#!/bin/bash

go install github.com/cosmtrek/air@v1.44.0
go install github.com/go-delve/delve/cmd/dlv@v1.20.2

make clean
make debug_api

air -c .air.toml
