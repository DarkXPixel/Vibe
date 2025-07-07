FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd/main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/auth-service /auth-service

ENTRYPOINT [ "/auth-service" ]