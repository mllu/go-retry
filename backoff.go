package retry

import "time"

// BackoffPolicy is an interface to define backoff policy
type BackoffPolicy interface {
	TimeToWait(n uint) time.Duration
}

// DefaultBackoffPolicy is default backoff policy
type DefaultBackoffPolicy struct {
	Interval    time.Duration // Time interval to increase exponentially
	MinWaitTime time.Duration // Minimum wait time
	MaxWaitTime time.Duration // Maximum wait time
	MaxRetries  uint          // Maximum number of retries
}

// NewDefaultBackoffPolicy initiates DefaultBackoff with default values
func NewDefaultBackoffPolicy(interval time.Duration) *DefaultBackoffPolicy {
	return &DefaultBackoffPolicy{
		Interval:    interval,
		MinWaitTime: 1 * interval,
		MaxWaitTime: 30 * interval,
		MaxRetries:  5,
	}
}

// TimeToWait increase numRetries and return the time to wait until the next retry
func (dbp *DefaultBackoffPolicy) TimeToWait(n uint) time.Duration {
	duration := (1 << n) * dbp.Interval
	return duration
}
