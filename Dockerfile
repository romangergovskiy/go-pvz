FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY . .  

RUN go mod tidy 
RUN go build -o main ./cmd

CMD ["./main"]
