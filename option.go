package retry

import "time"

// DefaultBackoffPolicyOption is functional options for default backoff policy
type DefaultBackoffPolicyOption func(*DefaultBackoffPolicy)

// Interval set the initial interval for retry and cascade to MinWaitTime and MaxWaitTime
// default is 1 second
func Interval(interval time.Duration) DefaultBackoffPolicyOption {
	return func(dbp *DefaultBackoffPolicy) {
		dbp.Interval = interval
		dbp.MinWaitTime = 1 * interval
		dbp.MaxWaitTime = 30 * interval
	}
}

// MinWaitTime set the minimjm wait time for retry, default is 1 interval
func MinWaitTime(minWaitTime time.Duration) DefaultBackoffPolicyOption {
	return func(dbp *DefaultBackoffPolicy) {
		dbp.MinWaitTime = minWaitTime
	}
}

// MaxWaitTime set the maximum wait time for retry, default is 1 interval
func MaxWaitTime(maxWaitTime time.Duration) DefaultBackoffPolicyOption {
	return func(dbp *DefaultBackoffPolicy) {
		dbp.MaxWaitTime = maxWaitTime
	}
}

// MaxRetries set the maximum for retries, default is 5
func MaxRetries(maxRetries uint) DefaultBackoffPolicyOption {
	return func(dbp *DefaultBackoffPolicy) {
		dbp.MaxRetries = maxRetries
	}
}
