#!/bin/bash

go install github.com/cosmtrek/air@v1.44.0

air -c .air.toml
