FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /my-go-app ./cmd/app

FROM alpine:latest

WORKDIR /

COPY --from=builder /my-go-app /my-go-app

EXPOSE 8080

ENTRYPOINT ["/my-go-app"]