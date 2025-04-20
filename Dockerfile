# Указываем образ с Go
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod tidy

# Копируем весь проект в контейнер
COPY . .

# Собираем проект, указывая путь к main.go
RUN go build -o main ./cmd

# Указываем команду, которая будет выполняться при старте контейнера
CMD ["./main"]
