FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine

WORKDIR /root/
COPY --from=builder /app/app .

CMD ["./app"]