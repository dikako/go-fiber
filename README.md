# Go Fiber Web Server

A high-performance web server built with Go Fiber framework, featuring middleware support and configurable timeouts.

## Features

- Fast HTTP server using Fiber framework
- Configurable idle, read, and write timeouts
- Middleware support with prefix routing
- Pre-fork mode enabled for better performance
- Simple API endpoints demonstration

## Prerequisites

- Go 1.24 or higher
- Git (for cloning the repository)

## Installation

1. Clone the repository:
```bash
git clone <your-repository-url>
cd go-fiber
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The server is configured with the following default settings:
- Idle Timeout: 5 seconds
- Read Timeout: 5 seconds
- Write Timeout: 5 seconds
- Pre-fork mode: Enabled
- Server Address: localhost:3000

## API Endpoints

### 1. Root Endpoint
- Path: `/`
- Method: `GET`
- Response: "Hello, World!"

### 2. API Version 1 Endpoint
- Path: `/api/v1`
- Method: `GET`
- Response: "Hello, World!"

## Middleware

The application includes middleware for the `/api` prefix that logs before and after each request.

## Running the Application

To start the server:

```bash
go run main.go
```

The server will start on `localhost:3000`.

## Testing

Run the tests using:

```bash
go test -v
```

## Dependencies

- github.com/gofiber/fiber/v2 - Web framework
- github.com/gofiber/template - Template engine support
- github.com/gofiber/template/mustache/v2 - Mustache template engine
- Other supporting packages (see go.mod for full list)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 