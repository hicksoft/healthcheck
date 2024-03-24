# healthcheck

A simple, lightweight container that will ping a healthcheck endpoint while a target service is healthy and stop pinging when the target service is not.

## Key Concepts

**Target**: A service that you want to monitor

**Ping**: An endpoint that will be pinged while **Target** is healthy

**Period** How frequently to check **Target**'s health

**Status** HTTP status code that signifies **Target** is healthy


Every **Period**, healthcheck will send a GET request to **Target**. If a HTTP status code equal to **Status** is returned, a GET request will be sent to **Ping**, an uptime monitor such as healthcheck.io (no affiliation). If **Status** is not returned, no request will be sent to **Ping**.

## Docker Compose

```
version: "3.8"
services:
  healthcheck:
    image: ghcr.io/s-wasser/healthcheck
    container_name: healthcheck
    volumes:
      - ./config:/config
    environment:
      - TZ=America/Toronto
    restart: unless-stopped
```

Monitor conigruations are defined in **/config/config.yml**

## Config
```
# All options
Monitor1:                           # Name of the monitor
  target: http://myservice2.com     # Required
  ping: http://healthceck.abc       # Required
  period: 30s                       # Optional. Default: 10m. Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
  status: 201                       # Optional. Default: 200

# Minimum options
Monitor2:
  target: http://myservice1.com
  ping: http://healthceck.xyz
```
