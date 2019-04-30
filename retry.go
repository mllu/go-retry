package retry

import (
	"context"
	"time"
)

// Function will be called when timer went off
type Function func() error

// Do will retry function f with default BackOff policy
func Do(ctx context.Context, f Function, opts ...DefaultBackoffPolicyOption) (uint, error) {
	dbp := NewDefaultBackoffPolicy(1 * time.Second)
	// configure option
	for _, opt := range opts {
		opt(dbp)
	}
	n := uint(0)
	var err error
	for ; n < dbp.MaxRetries; n++ {
		err = f()
		if err == nil {
			return n, nil
		}
		waitTime := dbp.TimeToWait(n)
		select {
		case <-time.After(waitTime):
		case <-ctx.Done():
			return n, err
		}
	}
	return n, err
}

// Backoff will retry function f with given BackOff policy
func Backoff(ctx context.Context, f Function, bp BackoffPolicy) (uint, error) {
	return 0, nil
}
