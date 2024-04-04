# Start with a base image containing the Go runtime
FROM golang:1.21.5-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM golang:1.21.5-bullseye

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

ARG PORT
EXPOSE $PORT

CMD ["./main"]
