package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"time"

	httpclient "github.com/gretro/webhook-fwd/src/http_client"
	"github.com/gretro/webhook-fwd/src/utils"
)

type Client struct {
	options    ClientOptions
	httpClient *http.Client
}

const (
	DefaultTimeout = 10 * time.Second
)

type ClientAuthorization struct {
	ApiToken string
}

type ClientOptions struct {
	Authorization *ClientAuthorization
	ServerUrl     string
	Agent         string
	RetryCount    int
	RetryDelay    time.Duration
}

func NewClient(options ClientOptions) *Client {
	if options.Agent == "" {
		options.Agent = "unknown"
	}

	return &Client{
		options: options,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func (c *Client) Channel(name string) *ChannelRef {
	return NewChannelRef(name, c)
}

func (c *Client) buildUrl(urlPath string) (string, error) {
	result, err := url.JoinPath(c.options.ServerUrl, urlPath)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) createReq(ctx context.Context, method string, path string) (*http.Request, error) {
	url, err := c.buildUrl(path)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http request; %w", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("%s %s (%s)", c.options.Agent, runtime.GOOS, runtime.GOARCH))

	if c.options.Authorization != nil && c.options.Authorization.ApiToken != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Basic :%s", c.options.Authorization.ApiToken))
	}

	return request, nil
}

func (c *Client) createReqWithBody(ctx context.Context, method string, path string, body interface{}) (*http.Request, error) {
	request, err := c.createReq(ctx, method, path)
	if err != nil {
		return nil, err
	}

	json, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling http request body: %w", err)
	}

	request.Body = io.NopCloser(bytes.NewReader(json))
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func (c *Client) performHttpRequest(req *http.Request, result interface{}) error {
	_, err := utils.Retry(func() (bool, error) {
		err := httpclient.PerformHttpRequest(c.httpClient, req, result)
		return false, err
	}, utils.RetryOptions{
		MaxAttempts: c.options.RetryCount,
		RetryDelay:  c.options.RetryDelay,
		ShouldRetry: httpclient.ShouldRetryHttpRequest,
	})

	return err
}
