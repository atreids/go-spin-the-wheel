#!/bin/sh

GOOS=linux GOARCH=arm go build -o spin_bot_arm ../main.go
