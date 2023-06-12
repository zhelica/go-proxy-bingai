FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -tags netgo -trimpath -o go-proxy-百晓通AI main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/go-proxy-百晓通AI .

EXPOSE 8080
CMD ["/app/go-proxy-百晓通AI"]