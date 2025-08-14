#!/bin/sh

GOOS=linux GOARCH=arm go build -o ../dist/spin_bot_arm ../app.go
