package utils

import (
	"fmt"
	"time"
)

type RetryOptions struct {
	MaxAttempts int
	RetryDelay  time.Duration
	ShouldRetry func(error) bool
}

const (
	MaxAttemptErrorText = "max attempt exceeded"
)

func Retry[T interface{}](fn func() (T, error), options RetryOptions) (T, error) {
	attempt := 0
	var result T
	var lastError error

	for attempt < options.MaxAttempts {
		result, lastError = fn()

		if lastError == nil {
			return result, nil
		}

		attempt++
		if attempt > options.MaxAttempts {
			return result, fmt.Errorf("%s: %w", MaxAttemptErrorText, lastError)
		}

		shouldRetry := options.ShouldRetry == nil || options.ShouldRetry(lastError)
		if !shouldRetry {
			return result, lastError
		}

		time.Sleep(options.RetryDelay)
	}

	return result, fmt.Errorf("%s: %w", MaxAttemptErrorText, lastError)
}
