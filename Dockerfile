# Build Stage
FROM golang:1.22 AS builder
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
RUN CGO_ENABLED=0 go build -v -o server

# Install python3 and curl in the build stage using apt-get
RUN apt-get update && \
    apt-get install -y python3 curl && \
    ln -sf python3 /usr/bin/python

# Install yt-dlp
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp

# Deploy Stage
FROM alpine:latest

# Install ffmpeg and Python in the final image
RUN apk add --no-cache ffmpeg python3

# Copy the built Go binary from the builder stage
COPY --from=builder /app/server .

# Copy yt-dlp from the builder stage
COPY --from=builder /usr/local/bin/yt-dlp /usr/local/bin/yt-dlp

# Define the entrypoint
CMD ["/server"]
