# ---------- Builder Stage ----------
FROM dhi.io/golang:1.26-debian13-sfw-ent-dev as builder
# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o sentinel .


# ---------- Runtime Stage ----------
FROM dhi.io/distroless/base:nonroot

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/sentinel .

# Expose port (if needed)
EXPOSE 8080

# Run as non-root (already nonroot in distroless)
USER nonroot

# Start application
CMD ["./sentinel"]
