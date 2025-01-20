# CUTU-2025-backend

## Stack
- golang
- go fiber
- postgres

## Getting Start
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- golang 1.22
- docker
- makefile

### Installing
1. Clone the repo
2. Copy `.env.example` to `.env` and fill the values
3. Run `go mod download` to download all the dependencies

### Running
1. Run `docker-compose up -d` to start the database (optional for local development database)
2. Run `make server` or `go run cmd/main.go` to start the server

### API
