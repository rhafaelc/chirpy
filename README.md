# Chirpy

Chirpy is a simple Go-based server application built as part of a learning exercise on [Boot.dev](https://boot.dev). It provides basic functionality for managing users, chirps (short messages), and simulates webhook handling. The project is designed to help you learn backend development concepts, including working with databases, REST APIs, and middleware.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Database Setup](#database-setup)
- [API Endpoints](#api-endpoints)
  - [Public Endpoints](#public-endpoints)
  - [Admin Endpoints](#admin-endpoints)
- [Static File Server](#static-file-server)
- [Acknowledgments](#acknowledgments)

## Features

- **User Management**: Create, login, update, and revoke user tokens.
- **Chirp Management**: Create, list, retrieve, and delete chirps.
- **Health Check**: Endpoint to check server readiness.
- **Metrics and Reset**: Admin endpoints for monitoring and resetting the server state.
- **Webhook Simulation**: Simulates handling webhooks from an external service.
- **Static File Server**: Serves static files from the `./app` directory.

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.24.0 or higher)
- [PostgreSQL](https://www.postgresql.org/download) (version 17.4 or higher)

Alternatively, you can use [devbox](https://www.jetpack.io/devbox) for development:

```sh
devbox shell

# Start PostgreSQL service
devbox services start postgresql
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/rhafaelc/chirpy.git
cd chirpy
```

2. Install dependencies:

```sh
go mod tidy
```

3. Set up `.env`:

```sh
cp .env.example .env
```

4. Run the application:

```sh
go run .

# or
go build -o out && ./out
```

## Database Setup

1.  Initialize a new PostgreSQL database cluster (if needed):

```sh
initdb
```

2.  Create a new PostgreSQL user:

```sh
createuser --interactive
```

3.  Create a database for Gator:

```sh
createdb chirpy
```

4.  Run migration:

```sh
# Install goose for migrations
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations (located in sql/schema)
goose up
```

## API Endpoints

### Public Endpoints

#### Health Check

`GET /api/healthz`
Returns the readiness status of the application.

#### User Management

`POST /api/users` - Create a new user.

`POST /api/login` - Login and retrieve a JWT token.

`POST /api/refresh` - Refresh a JWT token.

`POST /api/revoke` - Revoke a JWT token.

`PUT /api/users` - Update user details.

#### Chirp Management

`POST /api/chirps` - Create a new chirp.

`GET /api/chirps` - List all chirps.

`GET /api/chirps/{chirpID}` - Retrieve a chirp by ID.

`DELETE /api/chirps/{chirpID}` - Delete a chirp by ID.

#### Webhooks

`POST /api/polka/webhooks` - Simulates handling webhooks.

### Admin Endpoints

#### Metrics

`GET /admin/metrics` - Retrieve application metrics.

#### Reset

`POST /admin/reset` - Reset application state.

## Static File Server

Serves files from the `./app` directory.
Access via `http://localhost:8080/app/`.

## Acknowledgments

This project was built as part of the [Boot.dev](https://boot.dev) learning platform.
