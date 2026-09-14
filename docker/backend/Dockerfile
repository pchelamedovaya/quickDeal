FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .
RUN CGO_ENABLED=0 go build -o /out/quickdeal main.go

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /out/quickdeal .

EXPOSE 8080

CMD ["./quickdeal"]
