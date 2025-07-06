# Go E-Commerce API

A robust, modular e-commerce RESTful API built with Go. This project demonstrates clean architecture, layered design, JWT authentication, and integration with external services such as payment gateways. It is designed for learning, extensibility, and real-world use cases.

> **Note:** This project is implemented based on the [roadmap.sh eCommerce API project challenge](https://roadmap.sh/projects/ecommerce-api).

## Database Design Diagram

<p align="center"><img src='/docs/files/database_diagram.png' alt='Golang Web API System Design Diagram' /></p>

## Features

- **User Authentication:** Signup, login, JWT-based authentication, and role-based access control.
- **Product Management:** CRUD operations for products, categories, and product images.
- **Shopping Cart:** Add, update, and remove items from the cart.
- **Order Processing:** Place orders, manage order items, and track order status.
- **Payment Integration:** Connects to payment gateways (e.g., Zarinpal), handles payment verification, and updates order/payment status.
- **Product Reviews:** Users can leave reviews and ratings for products; average rating is auto-updated.
- **Admin Panel:** Admin endpoints for managing products, categories, and users.
- **Rate Limiting & Security:** Middleware for rate limiting, CORS, and authentication.
- **Clean Architecture:** Separation of concerns with domain, usecase, repository, and delivery layers.
- **Dockerized:** Ready for deployment with Docker and docker-compose.

## Project Structure

```
├── src/
│   ├── api/            # HTTP handlers, DTOs, routers
│   ├── cmd/            # Application entrypoint
│   ├── config/         # Configuration files
│   ├── constants/      # Constant values
│   ├── dependency/     # Dependency injection
│   ├── docs/           # Swagger docs
│   ├── domain/         # Models, business entities
│   ├── events/         # Domain events
│   ├── infra/          # Database, cache, payment gateway
│   ├── pkg/            # Shared packages (errors, etc.)
│   ├── usecase/        # Business logic
│   └── ...
├── docker/             # Docker and compose files
├── README.md           # Project documentation
└── ...
```

## Getting Started

### Prerequisites
- Go 1.20+
- Docker & Docker Compose (for local development)
- PostgreSQL & Redis (can use provided docker-compose)

### Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/alielmi98/go-ecommerce-api.git
   cd go-ecommerce-api
   ```
2. **Copy and edit config:**
   ```sh
   cp src/config/config-development.yml src/config/config.yml
   # Edit config.yml as needed
   ```
3. **Run with Docker Compose:**
   ```sh
   docker-compose up --build
   ```
   Or run locally:
   ```sh
   cd src
   go mod tidy
   go run cmd/main.go
   ```

### API Documentation
- Swagger UI available at `/docs/index.html` when the server is running.

## Usage
- Register and login to receive a JWT token.
- Use the token in the `Authorization` header for protected endpoints.
- Admin endpoints require an admin role.

## Key Endpoints
- `POST /api/v1/auth/signup` — User registration
- `POST /api/v1/auth/login` — User login
- `GET /api/v1/products` — List products
- `POST /api/v1/cart` — Add to cart
- `POST /api/v1/orders` — Place order
- `POST /api/v1/payments` — Create payment URL
- `POST /api/v1/payments/verify` — Verify payment
- `POST /api/v1/reviews` — Add product review

## Technologies Used
- **Go** (Golang)
- **Gin** (HTTP framework)
- **GORM** (ORM)
- **PostgreSQL** (database)
- **Redis** (cache, rate limiting)
- **Swagger** (API docs)
- **Docker**

## Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](LICENSE)

---

> **Note:** This project is for learning and demonstration purposes. For production use, review and enhance security, validation, and error handling as needed.
