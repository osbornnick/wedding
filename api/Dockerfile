FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./main.go

FROM alpine:latest
RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app /app
CMD ["/app/main"]