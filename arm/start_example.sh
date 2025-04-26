#!/bin/sh

# An example file for automatically starting the bot on the raspberry.

nohup ./spin_bot_arm64 >./logfile &
echo $! >./pid_file
