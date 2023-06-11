package utils_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gretro/webhook-fwd/utils"
	"github.com/stretchr/testify/suite"
)

type RetryUtilTestSuite struct {
	suite.Suite
}

func TestRetryUtils(t *testing.T) {
	suite.Run(t, new(RetryUtilTestSuite))
}

func (suite *RetryUtilTestSuite) Test_WhenSuccess_ShouldReturnResult() {
	data := "hello, world"
	fn := func() (string, error) {
		return data, nil
	}

	result, err := utils.Retry(fn, utils.RetryOptions{
		MaxAttempts: 5,
		RetryDelay:  100 * time.Millisecond,
		ShouldRetry: nil,
	})

	suite.NoError(err)
	suite.Equal(data, result)
}

func (suite *RetryUtilTestSuite) Test_WhenEventualSuccess_ShouldReturnResult() {
	data := "hello, world"
	attempt := 0

	fn := func() (string, error) {
		if attempt < 3 {
			attempt++
			return "", fmt.Errorf("error")
		}

		return data, nil
	}

	result, err := utils.Retry(fn, utils.RetryOptions{
		MaxAttempts: 5,
		RetryDelay:  100 * time.Millisecond,
		ShouldRetry: nil,
	})

	suite.NoError(err)
	suite.Equal(attempt, 3)
	suite.Equal(data, result)
}

func (suite *RetryUtilTestSuite) Test_WhenError_AndShouldNotRetry_ShouldReturnError() {
	attempts := 0
	fn := func() (string, error) {
		attempts++
		return "", errors.New("my test error")
	}

	result, err := utils.Retry(fn, utils.RetryOptions{
		MaxAttempts: 5,
		// We never retry here
		ShouldRetry: func(err error) bool { return false },
	})

	suite.Empty(result)
	suite.Error(err)
	suite.ErrorContains(err, "my test error")
	suite.Equal(attempts, 1)
}

func (suite *RetryUtilTestSuite) Test_WhenError_AndExceedsMaxRetries_ShouldReturnError() {
	attempts := 0
	maxAttempts := 10
	fn := func() (string, error) {
		attempts++
		return "", errors.New("my test error 2")
	}

	result, err := utils.Retry(fn, utils.RetryOptions{
		MaxAttempts: maxAttempts,
		RetryDelay:  10 * time.Millisecond,
		ShouldRetry: nil,
	})

	suite.Empty(result)
	suite.Error(err)
	suite.ErrorContains(err, "my test error 2")
	suite.ErrorContains(err, utils.MaxAttemptErrorText)
	suite.Equal(attempts, maxAttempts)
}
