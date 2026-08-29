# URL Shortener Analytics Dashboard

A companion gRPC microservice for the [URL Shortener](https://github.com/erfangho/url-shortener) project. Receives click events in real-time via gRPC and serves as the analytics backend.

## Relationship

```
┌─────────────────┐     gRPC      ┌──────────────────────┐
│  URL Shortener  │ ───────────→  │  This Service        │
│  (Client)       │  RecordClick  │  (gRPC Server)       │
│                 │               │  Port: 50051         │
└─────────────────┘               └──────────────────────┘
```

The URL Shortener project calls this service's `RecordClick` RPC every time a shortened URL is clicked.

## Tech Stack

- **Go 1.27**
- **gRPC** — Google's high-performance RPC framework
- **Protocol Buffers** — Service contract definition
- **Docker** — Multi-stage build with BuildKit cache mounts

## Project Structure

```
url-shortener-analytics/
├── cmd/server/          → gRPC server entrypoint
├── internal/grpc/       → Server implementation
├── proto/               → Protocol Buffer definitions + generated code
├── Dockerfile           → Multi-stage Docker build
└── dockerignore
```

## Setup

### Local

```bash
go run cmd/server/main.go
```

### Docker

```bash
docker build -t analytics-dashboard .
docker run -p 50051:50051 analytics-dashboard
```

## gRPC Service

### RecordClick

```proto
service AnalyticsService {
  rpc RecordClick(ClickEvent) returns (RecordResponse);
}
```

**Request:**
| Field | Type | Description |
|---|---|---|
| url_id | uint32 | ID of the shortened URL |
| user_agent | string | Client's User-Agent header |
| ip_address | string | Client's IP address |

**Response:**
| Field | Type | Description |
|---|---|---|
| success | bool | Whether the event was recorded |
| message | string | Status message |

### Test with grpcurl

```bash
# List available services
grpcurl -plaintext localhost:50051 list

# Record a click
grpcurl -plaintext -d '{"url_id": 1, "user_agent": "test", "ip_address": "127.0.0.1"}' \
  localhost:50051 analytics.AnalyticsService/RecordClick
```

## gRPC Concepts Demonstrated

| Concept | Implementation |
|---|---|
| Proto files | Service contract with message definitions |
| Code generation | `protoc` with Go plugins |
| gRPC Server | Implements generated `AnalyticsServiceServer` interface |
| Reflection | Service discovery for debugging tools |
| Client Wrapper | Hides proto types behind domain models in the caller |

## Environment Variables

| Variable | Description | Default |
|---|---|---|
| Port | gRPC server listen address | `:50051` |

## License

MIT
