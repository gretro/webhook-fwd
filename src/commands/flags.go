package commands

import "time"

const (
	VerboseFlag = "verbose"
	DebugFlag   = "debug"
	QuietFlag   = "quiet"
	ChannelFlag = "channel"
	ServerFlag  = "server"
	RetryFlag   = "retry"

	DefaultServer        = "http://localhost:25333"
	DefaultRetry         = 5
	DefaultRetryDuration = 500 * time.Millisecond
)
