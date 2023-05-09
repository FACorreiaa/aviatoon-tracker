FROM golang:1.20.4-alpine AS builder
LABEL stage=build

ENV CGO_ENABLED 0

ENV GOOS linux
RUN apk update --no-cache
WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["./main"]
