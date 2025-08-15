FROM golang:1.24.2 AS build

COPY . .
RUN GARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /wheel-bot ./app.go

FROM golang:1.24.2-alpine AS final

COPY --from=build /wheel-bot /usr/bin/

WORKDIR /usr/bin
RUN addgroup --gid 9999 wheel-bot && adduser -u 9999 -G wheel-bot -s /sbin/nologin --disabled-password wheel-bot
RUN chown wheel-bot:wheel-bot ./wheel-bot
USER wheel-bot
CMD [ "./wheel-bot" ]
