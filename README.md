Here's your fully enhanced `README.md`, now including:

* ✅ **Paseto authentication** with header usage
* ✅ **User & task endpoints**, request types, formats
* ✅ **Custom error handling**
* ✅ **Migration & mock data**
* ✅ **Detailed project structure explanation**

---

````markdown
# 📝 Task Note API

A secure, RESTful API for managing tasks, with full authentication using **Paseto**, structured via **Clean Architecture** in **Go (Gin)**. Supports file uploads, time zone–aware timestamps, custom errors, and Swagger-based documentation.

---

## 🔐 Authentication & Authorization

All **task-related endpoints** are protected with Paseto.

1. **Register** → `/api/v1/users/register`
2. **Login** → `/api/v1/users/login`

Login returns a token:

```json
{
  "access_token": "v2.local.eyJ..."
}
````

Use in requests:

```
Authorization: Bearer {token}
```

---

## 👤 User Endpoints

| Method | Endpoint                 | Description         | Request Body (JSON)                        |
|--------|--------------------------|---------------------|--------------------------------------------|
| POST   | `/api/v1/users/register` | Register a new user | `email`, `password`, `first_name`, `last_name` |
| POST   | `/api/v1/users/login`    | Login and get token | `email`, `password`                        |

---

## ✅ Task Features

* 🔒 Requires Paseto-based authentication
* ✅ All timestamps follow **RFC3339** in **Asia/Bangkok (+07:00)**
* ✅ Image upload via `multipart/form-data`
* ✅ Custom error response format with `code`, `message`, and optional `details`

---

## 📌 Task Endpoints

| Method | Endpoint            | Description          | Format                | Notes                                                         |
|--------|---------------------|----------------------|------------------------|---------------------------------------------------------------|
| POST   | `/api/v1/tasks`     | Create a new task     | `multipart/form-data` | Requires authentication. See form fields below               |
| PUT    | `/api/v1/tasks/:id` | Update a task by ID   | `multipart/form-data` | Only sends fields to update                                   |
| GET    | `/api/v1/tasks`     | Get list of tasks     | Query params          | Filterable & paginated                                        |
| GET    | `/api/v1/tasks/:id` | Get task by ID        | Path param            | Returns full task object                                      |
| DELETE | `/api/v1/tasks/:id` | Delete task by ID     | Path param            | Requires task ownership                                       |

---

## 🧾 Task Fields

### 🔸 Form Data Fields

#### For **Create** Task

| Field         | Type   | Required | Description                          |
|---------------|--------|----------|--------------------------------------|
| `title`       | string | ✅        | Max 100 characters                   |
| `status`      | string | ✅        | Must be `IN_PROGRESS` or `COMPLETED` |
| `date`        | string | ✅        | Format: `2025-05-04T14:30:00+07:00`  |
| `description` | string | ❌        | Optional                             |
| `image`       | file   | ❌        | Base64-encoded on backend            |

#### For **Update** Task

All fields are **optional** — only existing values will be updated:

| Field         | Type   | Required | Description                          |
|---------------|--------|----------|--------------------------------------|
| `title`       | string | ❌        | Max 100 characters                   |
| `status`      | string | ❌        | Must be `IN_PROGRESS` or `COMPLETED` |
| `date`        | string | ❌        | Format: `2025-05-04T14:30:00+07:00`  |
| `description` | string | ❌        | Optional                             |
| `image`       | file   | ❌        | Base64-encoded on backend            |

---

## 🔎 Query Parameters for `GET /api/v1/tasks`

| Param     | Type   | Required | Example      | Description                            |
|-----------|--------|----------|--------------|----------------------------------------|
| `search`  | string | ❌        | `Meeting`    | Search by title or description         |
| `sort_by` | string | ❌        | `created_at` | One of: `title`, `status`, `created_at` |
| `order`   | string | ❌        | `asc`        | Sort direction: `asc` or `desc`        |
| `limit`   | int    | ✅        | `10`         | Number of results per page (1–100)     |
| `offset`  | int    | ✅        | `1`          | Page number (starting at 1)            |

---

## 🔗 Path Parameters

These endpoints require a task ID in the path:

- `GET /api/v1/tasks/:id`
- `PUT /api/v1/tasks/:id`
- `DELETE /api/v1/tasks/:id`

Example:

---

## 🧱 Tech Stack

* Language: **Go (Golang)**
* Web framework: **Gin Gonic**
* Authentication: **Paseto (v2.local)**
* ORM: **GORM**
* Documentation: **Swagger via Swaggo**
* Testing: **Testify**
* Deployment: **Docker Compose**
* Timezone: **Asia/Bangkok (UTC+7)**
* Custom error response: `{ code, message, details? }`

---

## 📁 Project Structure Overview

```
.
├── config/         # Configuration
├── constant/       # Error codes, enums and constants
├── containers/     # Dependency injection wiring
├── controllers/    # HTTP request handlers
├── docs/           # Swagger documentation
├── entities/       # DTOs for request and response
├── infras/         # Logging, database, routes and server
├── middleware/     # Auth middleware (Paseto)
├── migration/      # DB migration + seed mock data
├── mocks/          # Mocks for testing
├── models/         # GORM models (Task, User)
├── repositories/   # Database access logic
├── services/       # Business logic layer
├── utils/          # Helpers: validation, time formatting, password, token
```

> 🔐 Token is generated using `PasetoMaker` under `utils/paseto.go`.

---

## 🚀 Getting Started

### Run with Docker

```bash
make reset
```

Swagger UI:

```
http://localhost:8080/api/v1/docs/index.html#/
```

---

## 🧪 Run Unit Tests

```bash
make test
```

Tests are written with **testify**, mocking all external dependencies.

---

## 🧾 Error Response Example

```json
{
  "error": {
    "code": 2001,
    "message": "Invalid request body",
    "details": [
      { "field": "title", "message": "Title is required" }
    ]
  }
}
```

---
