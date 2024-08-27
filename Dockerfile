FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/api/main.go


#start a new stage from scratch
FROM alpine:latest

# Install any necessary packages (optional, depending on the binary's requirements)
RUN apk add --no-cache libc6-compat

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main ./

EXPOSE 8080

# Run the binary
CMD [ "./main" ]

