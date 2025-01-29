# CryptoFolio

A Golang project for managing cryptocurrency portfolios. This project allows users to track the value of their cryptocurrency holdings by integrating with the CoinGecko API and check on benefits.

## Features

- **Portfolio Management**: Users can add, view, and manage their cryptocurrency holdings.
- **Almost Real-Time Pricing**: cryptocurrency prices are updated from CoinGecko every 10 minutes.
- **Portfolio Valuation**: Calculates the total value of a user's portfolio.
- **Test Coverage**: Includes unit tests for math functionalities.

---

## Project Structure

```
backend/
├── cmd/
│   └── main.go            # Entry point for the application
├── src/
│   ├── core/               # Handlers for HTTP routes
│   │   ├── asset.go
│   │   ├── cache.go
│   │   ├── crypto.go
│   │   ├── price.go
│   │   └── router.go
│   ├── db/                # Database connection and setup
│   │   └── db.go
│   │   └── reset_db.go    # Reset database
│   ├── utils/          # Business logic for pricing and portfolio
|   |   ├── math.go
|   |   ├── math_test.go
├── go.mod                 # Go module file
├── go.sum                 # Dependencies checksum
└── schema.sql             # Schema to init database

frontend/
├── public/
├── src/
│   ├── components/
│   │   ├── Add.tsx               # Add new assets
│   │   ├── Autocomplete.tsx      # autocomplete for assets
│   │   ├── ListDetails.tsx       # list detailed investment
│   │   ├── ListSum.tsx           # list global investment
│   │   ├── Totals.tsx            # show total invested and actual value
│   ├── lib/
│   │   ├── utils.ts
│   ├── pages/
│   │   ├── Homepage.tsx
│   ├── services/
│   │   ├── api.ts
│   │   ├── format.tsx
│   ├── App.css
│   ├── App.tsx
│   ├── index.css
│   ├── main.tsx
│   ├── vite-env.d.ts

```

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Bellzebuth/go-crypto.git
cd go-crypto
```

### 2. Build docker image

```bash
docker-compose build
```

### 3. Run docker

```bash
docker-compose up
```

The application will be available at `http://localhost:5173` and the API will be available at `http://localhost:8080`

## Testing

This project includes unit tests for math functionalities. To run the tests:

```bash
go test ./...
```

---

## Future Improvements

- **Wallet connexion**: Can connect directly your wallet and his history.
- **Authentication**: Add user authentication for secure portfolio access.
- **More Currencies**: Support all currencies.
- **Docker Support**: Add Docker configuration for easier deployment.

---

## License

This project is licensed under the MIT License. See `LICENSE` for details.

---

## Author

[Etienne Tournier Rigaudy](https://github.com/Bellzebuth)
