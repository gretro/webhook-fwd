package main

import (
	"fmt"
	"os"

	"github.com/gretro/webhook-fwd/config"
	"github.com/gretro/webhook-fwd/libs"
	"github.com/gretro/webhook-fwd/webserver/web"
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

	web.BootstrapWebServer(cfg)
}
