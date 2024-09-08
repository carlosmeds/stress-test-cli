FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/server .
ENTRYPOINT ["./server", "stress"]