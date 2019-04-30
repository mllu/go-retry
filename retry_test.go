package retry

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	testcases := []struct {
		name          string
		err           error
		interval      time.Duration
		expectRetries uint
	}{
		{
			name:          "no retry",
			err:           nil,
			interval:      10 * time.Millisecond,
			expectRetries: 0,
		},
		{
			name:          "default retry with 10 ms",
			err:           errors.New("test"),
			interval:      10 * time.Millisecond,
			expectRetries: 5,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			n, err := Do(
				context.Background(),
				func() error { return tc.err },
				Interval(tc.interval),
				MinWaitTime(1*tc.interval),
				MaxWaitTime(1*tc.interval),
				MaxRetries(5),
			)
			dur := time.Since(start)
			if shouldError := tc.err != nil; shouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			log.Println(dur)
			assert.Equal(t, tc.expectRetries, n)
		})
	}
}

func TestDoWithCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	start := time.Now()
	n, err := Do(
		ctx,
		func() error { return errors.New("test") },
	)
	dur := time.Since(start)
	assert.Error(t, err)
	log.Println(dur)
	assert.Equal(t, uint(1), n)
}

func TestBackoff(t *testing.T) {
	start := time.Now()
	n, err := Backoff(
		context.Background(),
		func() error { return nil },
		NewDefaultBackoffPolicy(10*time.Millisecond),
	)
	dur := time.Since(start)
	if shouldError := err != nil; shouldError {
		assert.Error(t, err)
	} else {
		assert.NoError(t, err)
	}
	log.Println(dur)
	assert.Equal(t, uint(0), n)
}
