FROM golang:1.24.2 AS build

COPY . .
RUN GARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /wheel-bot ./app.go

FROM golang:1.24.2-alpine AS final

COPY --from=build /wheel-bot /usr/bin/

WORKDIR /usr/bin
CMD [ "./wheel-bot" ]
