# Используем официальный образ Go
FROM golang:1.20 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
# Устанавливаем зависимости
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем исполняемый файл
RUN CGO_ENABLED=0 GOOS=linux go build -o task-service main.go

# Используем минимальный образ для выполнения
FROM alpine:latest

# Копируем исполняемый файл из этапа сборки
COPY --from=builder /app/task-service /task-service

# Открываем порт
EXPOSE 8082

# Запускаем приложение
CMD ["/task-service"]
