# Start from golang base image
FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Agus Wibawantara"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
#COPY /app/.env.dev .
#COPY configs .
COPY .env /app


# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENTRYPOINT go run main.go