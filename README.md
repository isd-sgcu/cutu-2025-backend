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
- air (optional for auto reload)

### Installing
1. Clone the repo
2. Copy `.env.example` to `.env` and fill the values
3. Run `go mod download` to download all the dependencies

### Running
#### Database
Run to start the local database for development
```sh
docker-compose up -d
```

#### Server
Option 1: Standard mode
```bash
make server
```
Option 2: Development mode with auto reload
```bash
air -c .air.toml
```

### API
