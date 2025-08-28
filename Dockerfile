# Build stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy only go.mod + go.sum first (so deps cache can be reused)
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .


# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .

# Copy environment variables file
COPY .env ./

# Copy the locales directory to the correct location
COPY ./locales ./locales

EXPOSE 8080

CMD ["./myapp"]
