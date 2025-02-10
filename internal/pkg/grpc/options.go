package grpc

import (
	"time"

	"google.golang.org/grpc/keepalive"
)

const (
	defaultTimeout = 60 * time.Second
)

type serverParameters struct {
	// MaxConnectionIdle is a duration for the amount of time after which an
	// idle connection would be closed by sending a GoAway. Idleness duration is
	// defined since the most recent time the number of outstanding RPCs became
	// zero or the connection establishment.
	maxConnectionIdle time.Duration // The current default value is infinity.
	// MaxConnectionAge is a duration for the maximum amount of time a
	// connection may exist before it will be closed by sending a GoAway. A
	// random jitter of +/-10% will be added to MaxConnectionAge to spread out
	// connection storms.
	maxConnectionAge time.Duration // The current default value is infinity.
	// MaxConnectionAgeGrace is an additive period after MaxConnectionAge after
	// which the connection will be forcibly closed.
	maxConnectionAgeGrace time.Duration // The current default value is infinity.
	// After a duration of this time if the server doesn't see any activity it
	// pings the client to see if the transport is still alive.
	// If set below 1s, a minimum value of 1s will be used instead.
	time time.Duration // The current default value is 2 hours.
	// After having pinged for keepalive check, the server waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	timeout time.Duration // The current default value is 20 seconds.
}

type Options func(*Service)

func (gs *Service) getKeepaliveParams() keepalive.ServerParameters {
	if gs.keepalive.timeout == 0 {
		gs.keepalive.timeout = defaultTimeout
	}

	return keepalive.ServerParameters{
		MaxConnectionIdle:     gs.keepalive.maxConnectionIdle,
		MaxConnectionAge:      gs.keepalive.maxConnectionAge,
		MaxConnectionAgeGrace: gs.keepalive.maxConnectionAgeGrace,
		Timeout:               gs.keepalive.timeout,
		Time:                  gs.keepalive.time,
	}
}
