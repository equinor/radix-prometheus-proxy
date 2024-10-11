package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"
)

type RouteMapper func(mux *http.ServeMux)

func NewRouter(handlers ...RouteMapper) *negroni.Negroni {
	mux := http.NewServeMux()
	for _, handler := range handlers {
		handler(mux)
	}

	return negroni.New(
		negroni.NewRecovery(),
		NewZerologRequestIdMiddleware(),
		NewLoggingMiddleware(),
		negroni.Wrap(mux),
	)
}

func Serve(ctx context.Context, port int, router http.Handler) error {

	s := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
	}
	go func() {

		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Ctx(ctx).Fatal().Msg(err.Error())
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	return s.Shutdown(shutdownCtx)
}
