package main

import (
	"fmt"
	"os"

	"github.com/gretro/webhook-fwd/src/config"
	"github.com/gretro/webhook-fwd/src/libs"
	"github.com/gretro/webhook-fwd/src/webserver"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("application panicked: %v\n", err)
			os.Exit(1)
		}
	}()

	cfg := config.BoostrapAppConfiguration()
	libs.BootstrapWebLogger(cfg)

	webserver.BootstrapWebServer(cfg)
}
