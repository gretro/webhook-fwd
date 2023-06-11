package httpclient

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gretro/webhook-fwd/src/libs"
	"go.uber.org/zap"
)

var retriableStatusCodes = []int{
	408,
	425,
	429,
	500,
	502,
	503,
	504,
}

func PerformHttpRequest(httpClient *http.Client, req *http.Request, result interface{}) error {
	logRequest(req)

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	logResponse(res, data)

	if res.StatusCode >= 400 {
		return &HttpError{
			statusCode: res.StatusCode,
		}
	}

	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

func ShouldRetryHttpRequest(err error) bool {
	httpError := HttpError{}
	isHttpError := errors.As(err, &httpError)

	if !isHttpError {
		return false
	}

	statusCode := httpError.StatusCode()

	for _, code := range retriableStatusCodes {
		if statusCode == code {
			return true
		}
	}

	return false
}

func logRequest(req *http.Request) {
	libs.Logger().Debug(
		"HTTP request",
		zap.String("url", req.URL.String()),
		zap.String("method", req.Method),
	)
}

func logResponse(res *http.Response, body []byte) {
	libs.Logger().Debug(
		"HTTP response",
		zap.String("url", res.Request.URL.String()),
		zap.String("method", res.Request.Method),
		zap.Int("status", res.StatusCode),
		zap.ByteString("response", body),
	)
}
