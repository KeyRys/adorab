## Adorab

Adorab is a rabbit marketplace backend built with Go. It provides RESTful APIs for authentication, rabbit listings, sellers, carts, checkout, orders, and rabbit recommendations.

The project is built using a Clean Architecture-inspired structure with PostgreSQL/Supabase as the database.

## Features

* JWT authentication
* User and seller management
* Rabbit/product management
* Rabbit image storage
* Shopping cart
* Checkout & order management
* Weighted Content-Based Filtering for rabbit recommendations
* Automated backend testing

## Tech Stack

* **Go**
* **Gin**
* **PostgreSQL**
* **Supabase**
* **JWT**
* **pgx**
* **Go Testing**

## Architecture

backend/
├── cmd/            # Application entry point
├── internal/
│   ├── delivery/   # HTTP handlers
│   ├── domain/     # Entities
│   ├── repository/ # Database operations
│   ├── route/      # API routes
│   └── usecase/    # Business logic
└── pkg/
    └── database/   # Database configuration

Request flow:

HTTP → Handler → Usecase → Repository → PostgreSQL

## 🧪 Testing

Automated backend tests are developed in the:
`test/backend-automated-testing`
branch.

Run tests with:
go test ./...

Verbose output:
go test -v ./...

Coverage:
go test ./... -cover

## 🚀 Getting Started

git clone https://github.com/KeyRys/adorab.git
cd adorab/backend
go mod download
go run ./cmd

Create a `.env` file with the required database, Supabase, and JWT configuration.

## 🌿 Branches

* **main** — Main backend development
* **test/backend-automated-testing** — Automated testing development

## Author

**KeyRys**

Computer Science graduate focused on backend and full-stack development.

GitHub: https://github.com/KeyRys
