# Crypto Portfolio Tracker

A Golang project for managing cryptocurrency portfolios. This project allows users to track the value of their cryptocurrency holdings in real-time by integrating with the CoinGecko API.

## Features

- **Portfolio Management**: Users can add, view, and manage their cryptocurrency holdings.
- **Real-Time Pricing**: Fetches real-time cryptocurrency prices from CoinGecko.
- **Portfolio Valuation**: Calculates the total value of a user's portfolio.
- **Charts**: Generates visualizations for portfolio performance over time.
- **Test Coverage**: Includes unit tests for key functionalities.

---

## Project Structure

```
backend/
├── cmd/
│   └── main.go            # Entry point for the application
├── src/
│   ├── api/               # Handlers for HTTP routes
│   │   ├── handlers.go
│   │   └── handlers_test.go
│   ├── db/                # Database connection and setup
│   │   └── db.go
│   ├── services/          # Business logic for pricing and portfolio
|   |   ├── charts.go
|   |   ├── pieCharts.go
│   │   ├── pricing.go
│   │   ├── portfolio.go
│   │   ├── pricing_test.go
│   │   └── portfolio_test.go
│   ├── test/
|   |   ├── integrations_test.go
├── go.mod                 # Go module file
├── go.sum                 # Dependencies checksum
└── README.md              # Project documentation
```

---

## Prerequisites

- **Golang**: Version 1.18 or higher
- **SQLite**: Used as the database for storing portfolio data

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Bellzebuth/go-crypto.git
cd go-crypto
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up the database

Run the following command to initialize the SQLite database:

```bash
sqlite3 portfolio.db < schema.sql
```

### 4. Run the application

Start the application:

```bash
go run cmd/main.go
```

The API will be available at `http://localhost:8080`.

---

## API Endpoints

### 1. Add a cryptocurrency to the portfolio

**POST** `/portfolio/add`

**Request Body:**

```json
{
  "user_id": 1,
  "name": "bitcoin",
  "amount": 0.1
}
```

**Response:**

```json
{
  "message": "Asset added successfully"
}
```

### 2. Get the total value of a portfolio

**GET** `/portfolio/value?user_id={user_id}`

**Response:**

```json
{
  "total_value_usd": 4500
}
```

---

## Testing

This project includes unit tests for key functionalities. To run the tests:

```bash
go test ./...
```

---

## Key Files

### **Handlers**

- `internal/api/handlers.go`: Defines HTTP endpoints for managing portfolios.
- `internal/api/handlers_test.go`: Unit tests for the handlers.

### **Services**

- `internal/services/pricing.go`: Fetches real-time cryptocurrency prices.
- `internal/services/portfolio.go`: Business logic for managing portfolios.
- `internal/services/pricing_test.go`: Unit tests for the pricing service.
- `internal/services/portfolio_test.go`: Unit tests for the portfolio service.

### **Database**

- `internal/db/db.go`: Contains database initialization and connection logic.

---

## Future Improvements

- **Authentication**: Add user authentication for secure portfolio access.
- **Enhanced Visualizations**: Provide detailed charts for portfolio trends.
- **Multiple Currencies**: Support valuation in currencies other than USD.
- **Docker Support**: Add Docker configuration for easier deployment.

---

## License

This project is licensed under the MIT License. See `LICENSE` for details.

---

## Author

[Etienne Tournier Rigaudy](https://github.com/Bellzebuth)
