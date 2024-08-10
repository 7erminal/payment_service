# Base image is in https://registry.hub.docker.com/_/golang/
# Refer to https://blog.golang.org/docker for usage
FROM golang:1.22-alpine AS builder
MAINTAINER Bede Abbe

# ENV GOPATH /go

WORKDIR /usr/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Install beego & bee
RUN go install github.com/beego/bee/v2@latest

# Build the Go app
RUN go build -o main .

# Build the Go app and create an executable named 'main'
RUN go build -o main .

# Stage 2: Create a smaller image to run the application
FROM alpine:latest

# Set the Current Working Directory inside the container to /app
WORKDIR /usr/app

# Copy the executable from the builder stage to /app
COPY --from=builder /usr/app/main .

# Copy configuration files (if any) from the builder stage to /app/conf
COPY --from=builder /usr/app/conf ./conf

# RUN bee run -downdoc=true -gendoc=true

EXPOSE 8088

CMD ["./main"]