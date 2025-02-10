package health

import (
	"net/http"
	"time"

	"github.com/alexliesenfeld/health"
	log "github.com/sirupsen/logrus"
)

func Logger() health.Middleware {
	return func(next health.MiddlewareFunc) health.MiddlewareFunc {
		return func(r *http.Request) health.CheckerResult {
			now := time.Now().UTC()
			result := next(r)
			log.WithFields(log.Fields{"healthceck_time": time.Now().Sub(now).String(), "result": result.Status}). // nolint
																Trace("[healthcheck] processed health check request")
			if result.Details != nil {
				for service, res := range result.Details {
					if res.Error != nil {
						log.WithField("err", res.Error).Errorf("[healthcheck] %s error", service)
					}
				}
			}
			return result
		}
	}
}
