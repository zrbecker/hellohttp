FROM golang:1.23-alpine AS builder

WORKDIR /hellohttp

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build ./...

FROM alpine:latest

WORKDIR /hellohttp

COPY --from=builder /hellohttp/build/server .

EXPOSE 8080
ENTRYPOINT ["./server"]
