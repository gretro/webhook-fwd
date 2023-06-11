package channels_test

import (
	"context"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/gretro/webhook-fwd/dtos"
	"github.com/gretro/webhook-fwd/errors"
	"github.com/gretro/webhook-fwd/webserver/channels"
	"github.com/stretchr/testify/suite"
)

type ChannelTestSuite struct {
	suite.Suite
}

func TestChannelSuite(t *testing.T) {
	suite.Run(t, new(ChannelTestSuite))
}

func (suite *ChannelTestSuite) Test_WhenRegisteringListenerTwice_ShouldFail() {
	channel := channels.NewChannel("test_channel")

	listenerName := "my-listener"
	listener := make(chan dtos.ChannelMessageDTO)

	err := channel.RegisterListener(listenerName, listener)
	if !suite.NoError(err) {
		return
	}

	err = channel.RegisterListener(listenerName, listener)
	if suite.Error(err) {
		suite.ErrorAs(err, &errors.AlreadyExistsError{})
	}
}

func (suite *ChannelTestSuite) Test_ShouldBroadcastToRegisteredListeners() {
	channel := channels.NewChannel("test_channel")

	ctx, done := context.WithCancel(context.Background())
	defer done()

	listener := make(chan dtos.ChannelMessageDTO)
	receivedMessages := make([]dtos.ChannelMessageDTO, 0)

	doneWg := sync.WaitGroup{}
	doneWg.Add(1)

	go func() {
		for {
			select {
			case message := <-listener:
				receivedMessages = append(receivedMessages, message)
			case <-ctx.Done():
				doneWg.Done()
				return
			}
		}
	}()

	err := channel.RegisterListener("test_listener", listener)
	if !suite.NoError(err) {
		return
	}

	message := dtos.ChannelMessageDTO{
		Channel:    channel.Name(),
		Headers:    map[string][]string{},
		Body:       []byte("test message"),
		ReceivedAt: time.Now(),
	}

	// Wait for the goroutine to be listening. There must be a better way...
	time.Sleep(100 * time.Millisecond)

	channel.Broadcast(message)

	done()
	doneWg.Wait()

	if suite.Len(receivedMessages, 1) {
		receivedMsg := receivedMessages[0]

		areEqual := reflect.DeepEqual(message, receivedMsg)
		suite.True(areEqual, "Objects differ. \nExpected: %+v\nReceived: %+v", message, receivedMsg)
	}
}

func (suite *ChannelTestSuite) Test_WhenListenerIsNotListening_ShouldNotBlock() {
	channel := channels.NewChannel("test_channel")

	listener := make(chan dtos.ChannelMessageDTO)

	err := channel.RegisterListener("test_listener", listener)
	if !suite.NoError(err) {
		return
	}

	message := dtos.ChannelMessageDTO{
		Channel:    channel.Name(),
		Headers:    map[string][]string{},
		Body:       []byte("my non-blocking test message"),
		ReceivedAt: time.Now(),
	}

	channel.Broadcast(message)

	// We are testing the Broadcast method returns immediately without blocking here. Thus, no assert is needed
}

func (suite *ChannelTestSuite) Test_WhenRemovingListener_ShouldCloseChannel() {
	channel := channels.NewChannel("remove_listener_channel")

	listenerName := "test_closing_listener"
	listener := make(chan dtos.ChannelMessageDTO)
	isListenerClosed := false

	ctx, done := context.WithCancel(context.Background())
	defer done()

	doneWg := sync.WaitGroup{}
	doneWg.Add(1)

	go func() {
		for {
			select {
			case _, hasMore := <-listener:
				isListenerClosed = !hasMore
			case <-ctx.Done():
				doneWg.Done()
				return
			}
		}
	}()

	err := channel.RegisterListener(listenerName, listener)
	if !suite.NoError(err) {
		return
	}

	// Wait for the goroutine to be listening. There must be a better way...
	time.Sleep(100 * time.Millisecond)

	err = channel.RemoveListener(listenerName)
	if !suite.NoError(err) {
		return
	}

	done()
	doneWg.Wait()

	suite.True(isListenerClosed, "listener should have closed")
}

func (suite *ChannelTestSuite) Test_WhenRemovingListenerThatWasNotRegistered_ShouldFail() {
	channel := channels.NewChannel("test-channel")

	err := channel.RemoveListener("test-listener")
	if suite.Error(err) {
		suite.ErrorAs(err, &errors.EntityNotFoundError{})
	}
}

func (suite *ChannelTestSuite) Test_WhenNoListener_ShouldReturnEmptyListenerNames() {
	channel := channels.NewChannel("test-channel")

	listeners := channel.GetListenerNames()
	suite.Empty(listeners)
}

func (suite *ChannelTestSuite) Test_ShouldReturnRegisteredListenerNames() {
	channel := channels.NewChannel("test-channel")

	listener := make(chan dtos.ChannelMessageDTO)

	listener1Name := "listener-1"
	listener2Name := "listener-2"
	listener3Name := "listerner-3"

	channel.RegisterListener(listener1Name, listener)
	channel.RegisterListener(listener2Name, listener)
	channel.RegisterListener(listener3Name, listener)

	listeners := channel.GetListenerNames()
	if suite.Len(listeners, 3) {
		suite.Contains(listeners, listener1Name)
		suite.Contains(listeners, listener2Name)
		suite.Contains(listeners, listener3Name)
	}

}
