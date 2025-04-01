# Go Fiber

A high-performance web server built with Go Fiber framework, featuring middleware support and configurable timeouts.

## Features

- Fast HTTP server using Fiber framework
- Configurable idle, read, and write timeouts
- Middleware support with prefix routing
- Pre-fork mode enabled for better performance
- Simple API endpoints demonstration
- Support for JSON, XML, and Form data parsing
- File upload capabilities
- Static file serving
- Template rendering with Mustache
- Comprehensive error handling

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
- Template Engine: Mustache

## API Endpoints

### Basic Endpoints
1. **Root Endpoint**
   - Path: `/`
   - Method: `GET`
   - Response: "Hello, World!"

2. **Hello with Query Parameter**
   - Path: `/hello`
   - Method: `GET`
   - Query Parameters: `name` (optional, defaults to "Guest")
   - Example: `/hello?name=Dika`
   - Response: "Hello {name}"

3. **Request Headers and Cookies**
   - Path: `/request`
   - Method: `GET`
   - Headers: `firstname`
   - Cookies: `lastname`
   - Response: "Hello {firstname} {lastname}"

### User and Order Management
4. **User Order Details**
   - Path: `/users/:userId/orders/:orderId`
   - Method: `GET`
   - URL Parameters:
     - `userId`: User identifier
     - `orderId`: Order identifier
   - Response: Order details for specific user

### Form Handling
5. **Form Submission**
   - Path: `/hello`
   - Method: `POST`
   - Content-Type: `application/x-www-form-urlencoded`
   - Form Parameters: `name`
   - Response: "Hello {name}"

6. **File Upload**
   - Path: `/upload`
   - Method: `POST`
   - Content-Type: `multipart/form-data`
   - Form Parameters: `file`
   - Response: Upload confirmation message
   - Note: Files are saved to ./target directory

### Authentication Endpoints
7. **Login**
   - Path: `/login`
   - Method: `POST`
   - Content-Type: `application/json`
   - Request Body:
     ```json
     {
       "username": "string",
       "password": "string"
     }
     ```
   - Response: "Hello {username}"

8. **Register**
   - Path: `/register`
   - Method: `POST`
   - Supports Multiple Content-Types:
     - `application/json`
     - `application/xml`
     - `application/x-www-form-urlencoded`
   - Request Body:
     ```json
     {
       "username": "string",
       "password": "string",
       "name": "string"
     }
     ```
   - Response: "Register {username} successfully"

### Static Files
- Serves static files from configured directory
- Supports file downloads

### Template Rendering
- Supports Mustache templates
- Templates located in ./template directory
- File extension: .mustache

## Error Handling

The application includes a custom error handler that returns:
- Status: 500 Internal Server Error
- Response: "Error: {error_message}"

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

The test suite includes comprehensive tests for all endpoints and features.

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