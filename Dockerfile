FROM golang:alpine AS builder
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/app ./cmd/main.go

FROM alpine:latest
WORKDIR /build

COPY --from=builder ../build/app .

CMD ["./app"]
