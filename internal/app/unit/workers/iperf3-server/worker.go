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
	for {
		select {
		case <-ctx.Done():
			logrus.Infof("%s stopped successfull", trace.FuncName())
			return nil
		default:
			if err := t.run(ctx); err != nil {
				logrus.WithError(trace.FuncNameWithError(err)).Errorf("worker")
			}
		}
	}
}

// nolint
func (t *Worker) run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "iperf3", "-s", "-p", t.conf.Iperf3Port)
	if err := cmd.Run(); err != nil {
		return trace.FuncNameWithErrorMsg(err, "running iperf")
	}
	return nil
}
