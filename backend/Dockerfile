# Stage 1: Build stage
FROM golang:1.16.4-alpine AS builder
LABEL stage=build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update --no-cache && apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/main ./main.go

# Stage 2: Production stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY backend/ /app/backend/

CMD ["./main"]
