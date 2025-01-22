Here's the updated `README.md` based on your new API routes:

---

# CUTU-2025 Backend

## Stack
- **Go** (Golang)
- **Fiber** (Go Web Framework)
- **PostgreSQL**

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.22 or later
- Docker
- Makefile
- **Air** (optional for auto-reload)

### Installing

1. **Clone the repository:**

   ```bash
   git clone https://github.com/isd-sgcu/cutu2025-backend.git
   cd cutu2025-backend
   ```

2. **Copy the environment configuration file:**

   ```bash
   cp .env.example .env
   ```

   Fill in the values in the `.env` file for your local environment.

3. **Download dependencies:**

   ```bash
   go mod download
   ```

### Running the Project

#### Database

To start the local database for development, run:

```bash
docker-compose up -d
```

This will launch the PostgreSQL database in a Docker container.

#### Server

Option 1: **Standard Mode**

To run the server normally:

```bash
make server
```

Option 2: **Development Mode (with auto-reload)**

To run the server in development mode with live auto-reload, use **Air** (a Go live reloading tool):

```bash
air -c .air.toml
```

This option will automatically reload the server when you change any Go files.

## User API Documentation

### Endpoints

#### GET /api/users

**Summary:**  
Get all users.

**Permission:** BearerAuth (Role: Staff or Admin)

**Description:**  
Retrieve a list of all users. Access is restricted to staff and admin roles.

**Response:**

| Status | Description               |
|--------|---------------------------|
| 200    | List of all users         |
| 500    | Failed to fetch users     |

#### GET /api/users/{id}

**Summary:**  
Get user by ID.

**Permission:** BearerAuth

**Description:**  
Retrieve a user by its ID.

**Request Parameters:**

| Parameter | Type   | Required | Description |
|-----------|--------|----------|-------------|
| id        | string | true     | User ID     |

**Response:**

| Status | Description              |
|--------|--------------------------|
| 200    | User details             |
| 404    | User not found           |
| 500    | Failed to fetch user     |

#### POST /api/users/register

**Summary:**  
Register a new user in the system.

**Permission:** No

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
| 500    | Failed to create user    |

**Success Response (201):**

```json
{
    "userId": "1345",
    "qrUrl": "http://localhost:4000/api/users/qr/1345",
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzc1MzMxNTQsIn..."
}
```

#### PATCH /api/users/{id}

**Summary:**  
Update user by ID.

**Permission:** BearerAuth (Role: Admin)

**Description:**  
Update a user by its ID.

**Request Parameters:**

| Parameter | Type   | Required | Description       |
|-----------|--------|----------|-------------------|
| id        | string | true     | User ID           |
| user      | object | true     | User data (JSON)  |

**Response:**

| Status | Description              |
|--------|--------------------------|
| 204    | User successfully updated|
| 400    | Invalid input            |
| 404    | User not found           |
| 500    | Failed to update user    |

#### PATCH /api/users/role/{id}

**Summary:**  
Update user role.

**Permission:** BearerAuth (Role: Admin)

**Description:**  
Update the role of a user (Admin, Staff, etc.).

**Request Parameters:**

| Parameter | Type   | Required | Description     |
|-----------|--------|----------|-----------------|
| id        | string | true     | User ID         |
| role      | string | true     | New Role        |

**Response:**

| Status | Description              |
|--------|--------------------------|
| 204    | User role successfully updated |
| 400    | Invalid role             |
| 404    | User not found           |
| 500    | Failed to update role    |

#### POST /api/users/qr/{id}

**Summary:**  
Scan QR code.

**Permission:** BearerAuth (Role: Staff or Admin)

**Description:**  
Scan a QR code associated with a user by their ID. This is accessible only for staff and admin roles.

**Response:**

| Status | Description              |
|--------|--------------------------|
| 200    | QR code scanned          |
| 404    | User not found           |
| 500    | Failed to scan QR code   |

#### GET /api/users/qr/{id}

**Summary:**  
Get QR code URL.

**Permission:** BearerAuth

**Description:**  
Retrieve the URL for a user's QR code by their ID.

**Response:**

| Status | Description              |
|--------|--------------------------|
| 200    | QR code URL              |
| 404    | User not found           |
| 500    | Failed to fetch user     |

---

This updated README includes all your routes, their descriptions, and the permissions required for each one. Let me know if you need further changes!