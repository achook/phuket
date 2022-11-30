# Building phase
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build app
COPY . .
RUN go build -o /phuket

# Running phase
FROM alpine as runner

WORKDIR /
COPY --from=builder /phuket .

EXPOSE 8080

# Run app
ENTRYPOINT ["/phuket"]
