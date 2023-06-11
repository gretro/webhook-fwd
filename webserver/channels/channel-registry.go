package channels

import (
	"sync"

	apperrors "github.com/gretro/webhook-fwd/errors"
)

type ChannelRegistry struct {
	channels map[string]*Channel

	rwMutex *sync.RWMutex
}

func NewChannelRegistry() *ChannelRegistry {
	return &ChannelRegistry{
		channels: make(map[string]*Channel, 0),
		rwMutex:  &sync.RWMutex{},
	}
}

func (reg *ChannelRegistry) GetChannel(name string) (*Channel, error) {
	reg.rwMutex.RLock()
	defer reg.rwMutex.RUnlock()

	channel, exists := reg.channels[name]
	if !exists {
		return nil, apperrors.NewEntityNotFoundError("Channel", name, nil)
	}

	return channel, nil
}

func (reg *ChannelRegistry) RegisterChannel(channel *Channel) error {
	reg.rwMutex.Lock()
	defer reg.rwMutex.Unlock()

	name := channel.Name()

	_, exists := reg.channels[name]
	if exists {
		return apperrors.NewAlreadyExistsError("Channel", name, nil)
	}

	reg.channels[name] = channel
	return nil
}

func (reg *ChannelRegistry) UnregisterChannel(name string) error {
	reg.rwMutex.Lock()
	defer reg.rwMutex.Unlock()

	_, exists := reg.channels[name]
	if !exists {
		return apperrors.NewEntityNotFoundError("Channel", name, nil)
	}

	delete(reg.channels, name)
	return nil
}
