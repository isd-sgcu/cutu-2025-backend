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

Here is the updated `README.md` based on the new API:


# User Management API

This API provides endpoints for managing users in the system, including retrieving, updating, and registering users. It also includes functionality for handling QR codes and managing user roles.

## API Endpoints

### 1. Get All Users

**Endpoint:** `/api/users`  
**Method:** `GET`

**Permission:** BearerAuth (Staff, Admin)

Retrieve a list of all users.

**Response:**
- `200 OK`: Returns a list of users.
- `500 Internal Server Error`: Failed to fetch users.

### 2. Get User by ID

**Endpoint:** `/api/users/{id}`  
**Method:** `GET`

**Permission:** BearerAuth

Retrieve a user by its ID.

**Parameters:**
- `id` (path) - The ID of the user.

**Response:**
- `200 OK`: Returns the user details.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to fetch user.

### 3. Register a New User

**Endpoint:** `/api/users/register`  
**Method:** `POST`

**Permission:** No


Register a new user in the system.

**Parameters (form data):**
- `id` (string) - User ID
- `name` (string) - User Name
- `email` (string) - User Email
- `phone` (string) - User Phone
- `university` (string) - User University
- `sizeJersey` (string) - Jersey Size
- `foodLimitation` (string) - Food Limitation
- `invitationCode` (string) - Invitation Code
- `status` (string) - User Status (`student`, `alumni`, `general_public`)
- `image` (file) - User Image
- `graduatedYear` (string) - Graduated Year
- `faculty` (string) - Faculty
- `education` (string) - User Education (`studying`, `graduated`)

**Response:**
- `201 Created`: User successfully created.
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized.
- `500 Internal Server Error`: Failed to create user.

### 4. Update User by ID

**Endpoint:** `/api/users/{id}`  
**Method:** `PATCH`

**Permission:** BearerAuth (Admin)

Update a user by its ID.

**Parameters:**
- `id` (path) - The ID of the user.
- `user` (body) - User data (JSON).

**Response:**
- `204 No Content`: User successfully updated.
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized.
- `403 Forbidden`: Forbidden.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to update user.

### 5. Get QR Code URL for User

**Endpoint:** `/api/users/qr/{id}`  
**Method:** `GET`

**Permission:** BearerAuth 

Retrieve a QR code URL for a user.

**Parameters:**
- `id` (path) - The ID of the user.

**Response:**
- `200 OK`: Returns the QR code URL.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to fetch user.

### 6. Scan QR Code

**Endpoint:** `/api/users/qr/{id}`  
**Method:** `POST`

**Permission:** BearerAuth (Staff, Admin)

Scan a QR code and perform associated actions.

**Parameters:**
- `id` (path) - The ID of the user.

**Response:**
- `200 OK`: User scanned successfully.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to process QR code.

### 7. Update User Role by ID

**Endpoint:** `/api/users/role/{id}`  
**Method:** `PATCH`

**Permission:** BearerAuth (Admin)

Update a user role by its ID.

**Parameters:**
- `id` (path) - The ID of the user.
- `role` (body) - User role (string).

**Response:**
- `204 No Content`: User role updated successfully.
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized.
- `403 Forbidden`: Forbidden.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to update user role.

### 8. Update Staff Info

**Endpoint:** `/api/users/`  
**Method:** `PATCH`

Update a staff role by its credentials.

**Parameters:**
- `user` (body) - User data (JSON).

**Permission:** BearerAuth

**Response:**
- `204 No Content`: User role updated successfully.
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized.
- `403 Forbidden`: Forbidden.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to update user role.

### 9. Remove Account

**Endpoint:** `/api/users/{id}`  
**Method:** `DELETE`

Update a staff role by its credentials.

**Parameters:**
- `id` (path) - The ID of the user.

**Permission:** BearerAuth(Admin)

**Response:**
- `204 No Content`: User role updated successfully.
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized.
- `403 Forbidden`: Forbidden.
- `404 Not Found`: User not found.
- `500 Internal Server Error`: Failed to update user role.

## Error Responses

### Error Response Format

```json
{
  "error": "Error message here"
}
```

### Common Error Codes
- `400 Bad Request`: Invalid input.
- `401 Unauthorized`: Unauthorized access.
- `403 Forbidden`: Forbidden action.
- `404 Not Found`: Resource not found.
- `500 Internal Server Error`: An error occurred on the server.

## Definitions

### Education Enum
- `studying`: The user is currently studying.
- `graduated`: The user has graduated.

### Role Enum
- `member`: A member user.
- `staff`: A staff user.
- `admin`: An admin user.

### Status Enum
- `student`: The user is a student.
- `alumni`: The user is an alumni.
- `general_public`: The user is from the general public.

### TokenResponse
- `accessToken`: The access token for authentication.
- `userId`: The user ID associated with the token.

### User
A user object containing:
- `id`: The user ID.
- `name`: The user name.
- `email`: The user email.
- `phone`: The user phone number.
- `status`: The user's status.
- `role`: The user's role.
- `education`: The user's education status.
- `imageURL`: The user's profile image URL.

---