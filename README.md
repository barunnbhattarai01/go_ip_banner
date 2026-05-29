# go_ip_banner

A small Go HTTP server that exposes a single `/` endpoint and wraps it with request rate limiting and temporary IP banning.

## What it does

- Starts an HTTP server on `0.0.0.0:8080`
- Serves `GET /`
- Returns a JSON response like:

  ```json
  {
    "message": "You are accessing data"
  }
  ```

- Limits requests per client IP using a token bucket limiter
- Bans an IP for 1 day if it exceeds the limit
- Cleans up inactive client entries periodically

## Requirements

- Go 1.25.4 or newer

## Run

From the project root:

```bash
go run ./cmd
```

Or build a binary:

```bash
go build -o bin/ip_banner ./cmd
./bin/ip_banner
```

Then open:

```text
http://localhost:8080/
```

## Behavior details

- The middleware reads the client IP from `RemoteAddr`
- The rate limiter is configured with a burst of 6 and a refill rate of 2 requests per second
- If a client exceeds the limit, it receives a JSON error response and is banned for 24 hours
- Inactive client state is removed after about 3 minutes

## Project layout

- `cmd/main.go` starts the server
- `cmd/handler.go` defines the response handler
- `cmd/middleware.go` contains rate limiting and IP banning logic
- `cmd/helpers.go` contains JSON response helpers
