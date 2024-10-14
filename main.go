package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os/signal"
	"time"

	prometheusApi "github.com/prometheus/client_golang/api"
	prometheusV1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/sys/unix"
)

var query = `min_over_time(probe_success{instance="https://api.dev.radix.equinor.com/health/"}[5m])`

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), unix.SIGTERM, unix.SIGINT)
	defer cancel()

	config := MustParseConfig()
	promController := NewPrometheusController(config)
	router := NewRouter(promController)

	log.Ctx(ctx).Info().Msgf("Starting server on http://localhost:%d/query...", config.Port)
	err := Serve(ctx, config.Port, router)
	log.Err(err).Msg("Terminated")
}

func NewPrometheusController(config Config) RouteMapper {
	apiClient, err := prometheusApi.NewClient(prometheusApi.Config{Address: config.Prometheus.String()})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create the Prometheus API client")
	}
	api := prometheusV1.NewAPI(apiClient)

	return func(mux *http.ServeMux) {
		mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
			logger := log.Ctx(r.Context())
			//
			end := time.Now()
			start := end.Add(24 * time.Hour * -30)
			promRange := prometheusV1.Range{Start: start, End: end, Step: 5 * time.Minute}
			content, warnings, err := api.QueryRange(r.Context(), query, promRange)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
				logger.Err(err).Msg("Failed to query prometheus")
				return
			}
			for _, w := range warnings {
				logger.Warn().Msg(w)
			}

			matrix, ok := content.(model.Matrix)
			if !ok {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
				logger.Error().Str("promtype", content.Type().String()).Type("type", content).Msg("Failed to parse response type")
				return
			}
			if len(matrix) != 1 {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
				logger.Error().Int("results", len(matrix)).Msg("the response did not have 1 item")
				return
			}

			body, err := json.Marshal(matrix[0].Values)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("Internal Server Error"))
				logger.Err(err).Msg("Failed to marshall json")
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(body)
		})
	}
}
