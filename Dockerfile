# Используем официальный образ Go для сборки приложения
FROM golang:1.22.5-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum и загружаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные исходные файлы
COPY . .

# Собираем Go-приложение
RUN go build -o /contractkeeper cmd/api/main.go

# Используем минимальный образ Alpine для запуска приложения
FROM alpine:latest

# Устанавливаем bash и необходимые зависимости
RUN apk --no-cache add bash

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /contractkeeper /contractkeeper

# Копируем wait-for-it скрипт и даем права на выполнение
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Копируем шаблоны и статические файлы
COPY templates /templates
COPY static /static

# Устанавливаем переменные окружения
ENV PORT=8080

# Открываем порт
EXPOSE 8080

# Добавляем команду для проверки состояния контейнера
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s \
  CMD wget --quiet --spider http://localhost:$PORT/health || exit 1

# Запускаем приложение с ожиданием базы данных
ENTRYPOINT ["/wait-for-it.sh", "db:5432", "--", "/contractkeeper"]