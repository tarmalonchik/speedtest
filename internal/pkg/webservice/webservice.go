package webservice

import (
	"context"
	"net/http"
	"time"

	"github.com/tarmalonchik/speedtest/internal/pkg/config"
	"github.com/tarmalonchik/speedtest/internal/pkg/trace"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type WebService struct {
	conf   config.BaseConfig
	router *mux.Router
}

func NewWebService(conf config.BaseConfig, router *mux.Router) *WebService {
	return &WebService{
		conf:   conf,
		router: router,
	}
}

func (s *WebService) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:         s.conf.GetAppAddr(),
		Handler:      s.router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	errC := make(chan error, 1)
	go func() {
		logrus.Infof("http servers listening %s", s.conf.GetAppAddr())
		errC <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		logrus.Info("stop http server")

		// nolint contextcheck here is parent context is closed. For timeout we should use context.Background()
		timeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		// nolint contextcheck
		err := server.Shutdown(timeout)
		if err != nil {
			logrus.WithError(trace.FuncNameWithError(err)).Errorf("failed to shutdown http server")
		}
		return nil
	case err := <-errC:
		return err
	}
}
