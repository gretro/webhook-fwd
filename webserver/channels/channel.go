package channels

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gretro/webhook-fwd/dtos"
	apperrors "github.com/gretro/webhook-fwd/errors"
	"github.com/gretro/webhook-fwd/libs"
	"go.uber.org/zap"
)

type ChannelListener = chan<- dtos.ChannelMessageDTO

type Channel struct {
	name      string
	createdAt time.Time

	rwMutex   *sync.RWMutex
	listeners map[string]ChannelListener
}

func NewChannel(name string) *Channel {
	return &Channel{
		name:      name,
		createdAt: time.Now(),

		rwMutex:   &sync.RWMutex{},
		listeners: make(map[string]ChannelListener, 0),
	}
}

func (c *Channel) Name() string {
	return c.name
}

func (c *Channel) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Channel) getListenerUnsafe(name string) (ChannelListener, error) {
	listener, exists := c.listeners[name]
	if !exists {
		return nil, apperrors.NewEntityNotFoundError("ChannelListener", name, nil)
	}

	return listener, nil
}

func (c *Channel) RegisterListener(name string, listener ChannelListener) error {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	_, err := c.getListenerUnsafe(name)
	if err == nil {
		return apperrors.NewAlreadyExistsError("ChannelListener", name, nil)
	}

	notFoundError := apperrors.EntityNotFoundError{}
	if !errors.As(err, &notFoundError) {
		return err
	}

	c.listeners[name] = listener
	return nil
}

func (c *Channel) RemoveListener(name string) error {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	listener, err := c.getListenerUnsafe(name)
	if err != nil {
		return err
	}

	close(listener)
	delete(c.listeners, name)

	return nil
}

func (c *Channel) Broadcast(message dtos.ChannelMessageDTO) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	logger := libs.Logger()

	for name, listener := range c.listeners {
		select {
		case listener <- message:
			logger.Debug("Message emitted", zap.String("listener", name))
		default:
			logger.Debug("Could not send message", zap.String("listener", name))
		}
	}
}

func (c *Channel) GetListenerNames() []string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	names := make([]string, len(c.listeners))
	i := 0
	for name := range c.listeners {
		names[i] = name
		i++
	}

	return names
}

func (c *Channel) String() string {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return fmt.Sprintf("channel %s (%d listeners)", c.name, len(c.listeners))
}
