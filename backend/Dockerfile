# Utiliser une image officielle de Go pour builder
FROM golang:1.23 as builder

# Définir le dossier de travail dans le conteneur
WORKDIR /app

# Copier les fichiers Go et les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le reste du code backend
COPY . .

# Compiler l'application Go
RUN go build -o main ./cmd/main.go

# Image légère pour exécuter l'application
FROM debian:bullseye-slim

# Définir le dossier de travail dans le conteneur
WORKDIR /app

# Copier le binaire depuis l'image builder
COPY --from=builder /app/main .
COPY schema.sql .

# Exposer le port utilisé par le backend
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./main"]
