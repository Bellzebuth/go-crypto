# load go image to build
FROM golang:1.23 as builder


WORKDIR /app

# copy go files and dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy backend code left
COPY . .

# compile go app
RUN go build -o main ./cmd/main.go

# light image to execute app
FROM debian:bookworm-slim


WORKDIR /app

# copy binary in builder
COPY --from=builder /app/main .
COPY schema.sql .

# install certificate
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*


# open port 8080
EXPOSE 8080

# execute app
CMD ["./main", "server"]

