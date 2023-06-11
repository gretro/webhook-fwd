package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gretro/webhook-fwd/src/dtos"
	apperrors "github.com/gretro/webhook-fwd/src/errors"
	httpclient "github.com/gretro/webhook-fwd/src/http_client"
)

const (
	ChannelEntityType = "Channel"
)

type ChannelRef struct {
	name   string
	client *Client
}

func NewChannelRef(name string, client *Client) *ChannelRef {
	return &ChannelRef{
		name:   name,
		client: client,
	}
}

func (c *ChannelRef) Name() string {
	return c.name
}

func (c *ChannelRef) Get(ctx context.Context) (dtos.ChannelDTO, error) {
	channelDto := dtos.ChannelDTO{}
	req, err := c.client.createReq(ctx, http.MethodGet, fmt.Sprintf("/channels/%s", c.name))
	if err != nil {
		return channelDto, err
	}

	err = c.client.performHttpRequest(req, &channelDto)
	if err != nil {
		return channelDto, fmt.Errorf("error getting channel: %w", err)
	}

	return channelDto, nil
}

func (c *ChannelRef) Exists(ctx context.Context) (bool, error) {
	_, err := c.Get(ctx)

	httpError := &httpclient.HttpError{}
	if errors.As(err, httpError) {
		if httpError.StatusCode() == http.StatusNotFound {
			return false, nil
		}
	}

	return err == nil, err
}

func (c *ChannelRef) Create(ctx context.Context) (dtos.ChannelDTO, error) {
	createChannelBody := dtos.CreateChannelDTO{
		Name: c.name,
	}

	channelDto := dtos.ChannelDTO{}

	req, err := c.client.createReqWithBody(ctx, http.MethodPost, "/channels", &createChannelBody)
	if err != nil {
		return channelDto, err
	}

	err = c.client.performHttpRequest(req, &channelDto)
	if err != nil {
		httpError := &httpclient.HttpError{}
		if errors.As(err, httpError) {
			if httpError.StatusCode() == http.StatusConflict {
				return channelDto, apperrors.NewAlreadyExistsError(ChannelEntityType, c.name, httpError)
			}
		}

		return channelDto, fmt.Errorf("error creating channel: %w", err)
	}

	return channelDto, nil
}

// TODO: Review return type
func (c *ChannelRef) Connect(ctx context.Context) error {
	return nil
}

func (c *ChannelRef) Close() error {
	return nil
}
