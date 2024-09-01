Certainly! Here's the README template in English:

---

# Go WebSocket Server

This is a WebSocket project written in Go using the Gorilla package, enabling real-time communication between a server and multiple clients.

## Key Features
- Real-time bidirectional communication.
- Support for multiple simultaneous clients.
- Connection and message management.

## Prerequisites

- [Go](https://golang.org/dl/) version 1.23.0 or higher
- [Gorilla WebSocket](https://github.com/gorilla/websocket)

## Installation

```bash
https://github.com/joaofilippe/go-websocket.git
cd go-webscoket
go mod tidy
```

## Usage

To start the WebSocket server, run:

```bash
go run main.go
```

You can use a WebSocket client, such as `websocat`, or Postman to test the connection.

## Project Structure

```
.
├── models
│   └── message.go
├── codealike.json
├── go.mod
├── go.sum
└── main.go
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/MyFeature`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature/MyFeature`).
5. Open a Pull Request.

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

- **Email**: joaofilippe@outlook.com
- **LinkedIn**: [João Filippe Rossi Rodrigues](https://www.linkedin.com/in/joaofilippe/)
