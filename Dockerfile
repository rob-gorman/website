# Build stage
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /docker-gs-ping

## Deploy
FROM scratch

LABEL org.opencontainers.image.source https://github.com/rob-gorman/webserver

COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 3000

ENTRYPOINT ["/docker-gs-ping"]
