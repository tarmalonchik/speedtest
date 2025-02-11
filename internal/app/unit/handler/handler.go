package handler

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/speedtest/internal/app/unit/svc"
)

func NewHandler(
	svc *svc.Service,
) *Handler {
	return &Handler{
		svc: svc,
	}
}

type Handler struct {
	svc *svc.Service
}

func (s *Handler) Speedtest(w http.ResponseWriter, r *http.Request) {
	outbound, inbound, err := s.svc.GetNodeSpeed(r.Context())
	if err != nil {
		logrus.Errorf("failed to get node speed: %v", err)
	}

	out := ""
	out += "# HELP speedtest_download_bytes Download Speed" + "\n"
	out += "# TYPE speedtest_download_bytes gauge" + "\n"
	out += fmt.Sprintf("speedtest_download_bytes %d", inbound) + "\n"
	out += "# HELP speedtest_upload_bytes Upload Speed" + "\n"
	out += "# TYPE speedtest_upload_bytes gauge" + "\n"
	out += fmt.Sprintf("speedtest_upload_bytes %d", outbound)

	_, _ = w.Write([]byte(out))
}
