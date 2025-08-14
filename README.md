# go-spin-the-wheel

A Go Discord bot which will give you a random Marvel Rival's hero to play or randomise the players in your voice channel into 2 separate teams.

## Commands

- `!spin`
- `!random_teams`

## Local dev

You need a `.env` file with `BOT_TOKEN=<bot-token>` in it.

Token can be regenerated from [Discord Developer Applications](https://discord.com/developers/applications) dashboard.

Build into exec:

> `go build app.go`

> `./app`

Run without building:

> `go run app.go`

## Running the bot on a Pi 

If you happen to have a Raspberry Pi (like me) you can find some helpful files in the `arm/` directory which will allow you to:

1. Cross-compile the bot for the Pi architecture.
2. An example systemd service to support using systemd to automatically start the bot on Pi boot.

### Systemd

To make a new service:

1. Create new systemd service file in `/etc/systemd/system/<service-name>.service` using the one in `arm/` as an example.
2. Reload daemon to include new file `sudo systemctl daemon-reload`.
3. Enable new service `sudo systemctl enable spin_bot.service`.
4. Either reboot system, or run `sudo systemctl start spin_bot.service` to start immediately.

Logs for your service can be viewed as normal using journalctl - `journalctl -u spin_bot.service`.
