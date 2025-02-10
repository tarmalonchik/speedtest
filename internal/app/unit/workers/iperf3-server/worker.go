package iperf3server

import (
	"context"
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/tarmalonchik/speedtest/internal/pkg/trace"
)

type Worker struct {
	conf Config
}

func NewWorker(conf Config) *Worker {
	return &Worker{
		conf: conf,
	}
}

func (t *Worker) Run(ctx context.Context) error {
	if t.conf.IsClient {
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		default:
			if err := t.run(ctx); err != nil {
				logrus.WithError(err).Errorf("worker")
			}
		}
	}
}

func (t *Worker) run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "iperf3", "-s")
	if err := cmd.Run(); err != nil {
		return trace.FuncNameWithErrorMsg(err, "running iperf")
	}
	return nil
}
