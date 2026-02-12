# ---------- Builder Stage ----------
FROM dhi.io/golang:1.26-debian13-sfw-ent-dev as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/sentinel ./cmd/sentinel

# ---------- Runtime Stage ----------
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app

COPY --from=builder /app/sentinel .

EXPOSE 8080
USER nonroot
CMD ["./sentinel"]
