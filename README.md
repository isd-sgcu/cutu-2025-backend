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

# User API Documentation

## Endpoints

### POST /api/users/register

**Summary:**
Register a new user in the system.

**Permission**: No

**Description:**
This endpoint allows users to register by submitting their personal information and an image.

**Request Format:** `multipart/form-data`

**Request Parameters:**

| Parameter      | Type   | Required | Description           |
|----------------|--------|----------|-----------------------|
| id             | string | true     | User ID               |
| name           | string | true     | User Name             |
| email          | string | true     | User Email            |
| phone          | string | true     | User Phone            |
| university     | string | true     | User University       |
| sizeJersey     | string | true     | Jersey Size           |
| foodLimitation | string | false    | Food Limitation       |
| invitationCode | string | false    | Invitation Code       |
| state          | string | true     | User State            |
| image          | file   | true     | User Image            |

**Response:**

| Status | Description              |
|--------|--------------------------|
| 201    | User successfully created|
| 400    | Invalid input            |
| 401    | Unauthorized             |
| 500    | Failed to create user    |

**Success Response (201):**
```json
{
    "userId": "1345",
    "qrUrl": "http://localhost:4000/api/users/qr/1345",
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc1MzMxNTQsInJvbGUiOiJzdGFmZiIsInVzZXJJZCI6IjEzNDUifQ.fdFfvVg--OdPqQK4iCSRx0PIY0IOlPjyojLjhvl9N4Q",
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzgxMzUyNTQsInJvbGUiOiJzdGFmZiIsInVzZXJJZCI6IjEzNDUifQ.HSyizL_f-0hetCaKtNtxTcDVu8zvfwpwFUanKI0n_Gw"
}
```

---

### GET /api/users

**Summary:**
Retrieve a list of all users.

**Permission**: Staff, Admin

**Response:**

| Status | Description              |
|--------|--------------------------|
| 200    | Successful response      |
| 500    | Failed to fetch users    |

---

### GET /api/users/{id}

**Summary:**
Retrieve a user by ID.

**Permission**: No

**Parameters:**

| Parameter | Type   | Required | Description |
|-----------|--------|----------|-------------|
| id        | string | true     | User ID     |

**Response:**

| Status | Description      |
|--------|------------------|
| 200    | User found       |
| 404    | User not found   |
| 500    | Failed to fetch user |

---

### PATCH /api/users/{id}

**Summary:**
Update a user by ID.

**Permission**: Admin

**Request Format:** `application/json`

**Parameters:**

| Parameter | Type   | Required | Description |
|-----------|--------|----------|-------------|
| id        | string | true     | User ID     |
| user      | object | true     | User data   |

**Response:**

| Status | Description      |
|--------|------------------|
| 204    | User updated     |
| 400    | Invalid input    |
| 401    | Unauthorized     |
| 403    | Forbidden        |
| 404    | User not found   |
| 500    | Failed to update user |

---

### GET /api/users/qr/{id}

**Summary:**
Retrieve user information by scanning QR code.

**Permission**: Staff, Admin

**Parameters:**

| Parameter | Type   | Required | Description |
|-----------|--------|----------|-------------|
| id        | string | true     | User ID     |

**Response:**

| Status | Description                        |
|--------|------------------------------------|
| 200    | Scan QR success                    |
| 400    | User has already entered           |
| 404    | User not found                     |
| 500    | Failed to process request          |

