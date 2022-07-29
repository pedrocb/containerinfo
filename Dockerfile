FROM golang:1.18.3-alpine as builder

WORKDIR /app

# Install dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy codebase
COPY *.go ./

# Build binary
RUN go build -o /out/containerinfo

FROM alpine:3.16.0

# Run as non privileged user
RUN adduser -H -D app
USER app

# Copy binary to lighter image
COPY --from=builder /out/containerinfo /containerinfo


EXPOSE 8000

CMD ["/containerinfo"]
