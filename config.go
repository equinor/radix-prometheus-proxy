package main

import (
	"net/url"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel    string `envconfig:"LOG_LEVEL" default:"info"`
	LogPretty   bool   `envconfig:"LOG_PRETTY" default:"false"`
	Port        int    `envconfig:"PORT" default:"8000"`
	QueriesFile string `envconfig:"QUERIES" default:"queries.yaml"`

	Prometheus url.URL `envconfig:"PROMETHEUS" required:"true"`
}

func MustParseConfig() Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		_ = envconfig.Usage("", &c)
		log.Fatal().Msg(err.Error())
	}

	initLogger(c)
	log.Info().Msg("Starting")
	log.Info().Int("Port", c.Port).Send()
	log.Info().Str("Log level", c.LogLevel).Send()
	log.Info().Bool("Log pretty", c.LogPretty).Send()
	log.Info().Stringer("Prometheus", &c.Prometheus).Send()
	log.Info().Str("Queries file", c.QueriesFile).Send()

	return c
}
