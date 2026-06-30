# SpotSync

A powerful synchronization solution for managing and syncing data across multiple sources efficiently.

## Features

- Real-time data synchronization
- Multi-source support
- Scalable architecture
- Error handling and retry logic
- Comprehensive logging
- RESTful API endpoints



## Installation

```bash
git clone https://github.com/Shohel-Raj/spotsync
cd spotsync
go mod download
```

## Quick Start

```bash
go run main.go
```

## Configuration

Configuration can be set via environment variables :

```yaml
server:
  PORT: 8080
  DSN: localhost
  JWT_SECRET: your_jwt_secret_key

```




### API Endpoints

- `GET /health` - Health check



# API Routes Documentation

Base URL:

```text
/api/v1
```

## Authentication

| Method | Endpoint | Access | Description |
|---------|----------|--------|-------------|
| POST | `/auth/register` | 🌐 Public | Register a new user account. |
| POST | `/auth/login` | 🌐 Public | Authenticate a user and return a JWT access token. |
| GET | `/auth/me` | 🔒 Private | Get the currently authenticated user's profile information. |

---

## Reservation Routes

| Method | Endpoint | Access | Description |
|---------|----------|--------|-------------|
| POST | `/reservations` | 🔒 Private | Create a new parking reservation for the authenticated user. |
| GET | `/reservations/my-reservations` | 🔒 Private | Retrieve all reservations created by the authenticated user. |
| GET | `/reservations` |  Admin Only | Retrieve all reservations from all users. |
| GET | `/reservations/:id` | 🔒 Private | Retrieve a specific reservation by its ID. |
| DELETE | `/reservations/:id` | 🔒 Private | Cancel a reservation by its ID. |

> **Note:**  
> `:id` represents the Reservation ID.

---

## Parking Zone Routes

| Method | Endpoint | Access | Description |
|---------|----------|--------|-------------|
| POST | `/parkingzones` | 👑 Admin Only | Create a new parking zone. |
| GET | `/parkingzones` | 🌐 Public | Retrieve all available parking zones. |
| GET | `/parkingzones/:id` | 🌐 Public | Retrieve a specific parking zone by its ID. |
| PUT | `/parkingzones/:id` | 👑 Admin Only | Update an existing parking zone. |
| DELETE | `/parkingzones/:id` | 👑 Admin Only | Delete a parking zone. |

> **Note:**  
> `:id` represents the Parking Zone ID.

---

# Access Levels

| Icon | Meaning |
|------|---------|
| 🌐 Public | No authentication required. |
| 🔒 Private | Requires a valid JWT access token. |
| 👑 Admin Only | Requires authentication and the user must have the **Admin** role. |

---

# Authentication

For every **Private** or **Admin Only** endpoint, include the JWT token in the request header.

```http
Authorization: Bearer <your_access_token>
```

Example:

```http
GET /api/v1/auth/me HTTP/1.1
Host: localhost:8080
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

# Route Summary

| Module | Public | Private | Admin |
|---------|:------:|:-------:|:-----:|
| Authentication | ✅ | ✅ | ❌ |
| Reservations | ❌ | ✅ | ✅ |
| Parking Zones | ✅ | ❌ | ✅ |

---

# Authorization Flow

```text
Public User
    │
    ├── Register
    └── Login
            │
            ▼
      Receive JWT Token
            │
            ▼
Authenticated User
            │
            ├── View My Profile
            ├── Create Reservation
            ├── View My Reservations
            ├── View Reservation Details
            └── Delete Reservation
            │
            ▼
        Admin User
            │
            ├── View All Reservations
            ├── Create Parking Zone
            ├── Update Parking Zone
            └── Delete Parking Zone
```




