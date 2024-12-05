# Go API Example

This is a simple example of a RESTful API written in Go.

## Getting Started

### Prerequisites

- Docker
- Git

### Installation

1. Clone the repository:
	```sh
	git clone https://github.com/yourusername/go-api-example.git
	```
2. Navigate to the project directory:
	```sh
	cd go-api-example
	```
3. Set file permissions:
	```sh
	sudo groupadd -g 1009 go-api
	sudo chgrp -R go-api ./
	sudo chmod g+w -R ./
	```
4. Install dependencies:
	```sh
	docker compose up -d
	```

The server will start on `http://localhost:8080`.

### API Endpoints

- `GET /receipts/{id}/points` - Retrieve the number of points earned for a receipt with the specified ID
- `POST /receipts/process` - Create a new receipt and process points

### License

This project is licensed under the GNU GPL v3 License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

- [Gin Gonic](https://github.com/gin-gonic/gin) - HTTP web framework for Go
