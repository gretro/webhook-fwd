package channels_test

import (
	"testing"

	"github.com/gretro/webhook-fwd/errors"
	"github.com/gretro/webhook-fwd/webserver/channels"
	"github.com/stretchr/testify/suite"
)

type ChannelRegistryTestSuite struct {
	suite.Suite
}

func TestChannelRegistrySuite(t *testing.T) {
	suite.Run(t, new(ChannelRegistryTestSuite))
}

func (suite *ChannelRegistryTestSuite) Test_WhenGetChannel_ShouldReturnChannel() {
	registry := channels.NewChannelRegistry()

	channel := channels.NewChannel("my-channel")
	err := registry.RegisterChannel(channel)

	if !suite.NoError(err) {
		return
	}

	result, err := registry.GetChannel(channel.Name())
	if !suite.NoError(err) {
		return
	}

	suite.Equal(channel, result)
}

func (suite *ChannelRegistryTestSuite) Test_WhenGetChannel_AndChannelNotFound_ShouldReturnError() {
	registry := channels.NewChannelRegistry()

	channel := channels.NewChannel("my-channel")

	result, err := registry.GetChannel(channel.Name())

	if suite.Error(err) {
		suite.ErrorAs(err, &errors.EntityNotFoundError{})
	}

	suite.Nil(result)
}

func (suite *ChannelRegistryTestSuite) Test_WhenRegisteringChannel_AndChannelAlreadyRegistered_ShouldReturnError() {
	registry := channels.NewChannelRegistry()

	channel := channels.NewChannel("my-channel")

	err := registry.RegisterChannel(channel)
	if !suite.NoError(err) {
		return
	}

	err = registry.RegisterChannel(channel)
	if suite.Error(err) {
		suite.ErrorAs(err, &errors.AlreadyExistsError{})
	}
}

func (suite *ChannelRegistryTestSuite) Test_WhenUnregisteringChannel_ShouldSucceed() {
	registry := channels.NewChannelRegistry()

	channel := channels.NewChannel("my-channel")

	err := registry.RegisterChannel(channel)
	if !suite.NoError(err) {
		return
	}

	err = registry.UnregisterChannel(channel.Name())
	suite.NoError(err)
}

func (suite *ChannelRegistryTestSuite) Test_WhenUnregistreringChannel_AndNotFound_ShouldReturnError() {
	registry := channels.NewChannelRegistry()

	err := registry.UnregisterChannel("my-channel")
	if suite.Error(err) {
		suite.ErrorAs(err, &errors.EntityNotFoundError{})
	}
}
