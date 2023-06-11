package httpclient

import "fmt"

type HttpError struct {
	statusCode int
}

func (err HttpError) Error() string {
	return fmt.Sprintf("http error: received status code %d", err.statusCode)
}

func (err HttpError) StatusCode() int {
	return err.StatusCode()
}
