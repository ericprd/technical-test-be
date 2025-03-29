package config

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chimware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func initRouter() chi.Router {
	router := chi.NewRouter()
	localCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "X-Device", "X-Client-Ip"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(
		localCors.Handler,
		chimware.RequestID,
	)

	return router
}

func NewRouter(apiRouter chi.Router) {
	mainRouter := chi.NewRouter()
	mainRouter.Mount("/api", apiRouter)
	mainCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: mainRouter,
	}

	group, groupCtx := errgroup.WithContext(mainCtx)

	group.Go(func() error {
		logrus.Infof("Listening to port %d", 8080)
		return svr.ListenAndServe()
	})

	group.Go(func() error {
		<-groupCtx.Done()

		ctxTimeout, cancel := context.WithTimeout(mainCtx, time.Duration(15)*time.Second)
		defer cancel()

		svr.Shutdown(ctxTimeout)

		return nil
	})

	if err := group.Wait(); err != nil {
		logrus.Errorf("system exit, reason: %v", err.Error())
	} else {
		logrus.Info("system exit normally")
	}
}
