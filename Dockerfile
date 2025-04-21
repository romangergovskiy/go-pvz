FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache git postgresql-client

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd

COPY wait-for-postgres.sh ./
RUN chmod +x wait-for-postgres.sh

CMD ["./wait-for-postgres.sh", "./main"]
