FROM golang:1.21.5-alpine

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all Go files
COPY . .

# Build the Go application
RUN go build -o /go-auth-app

EXPOSE 8080

CMD ["/go-auth-app"]