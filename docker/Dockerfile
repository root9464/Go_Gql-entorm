# Этап сборки
FROM golang:alpine AS builder

LABEL stage=gobuilder
ENV CGO_ENABLED=0 GOOS=linux
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

# Копируем зависимости для их установки
COPY ../go.mod ../go.sum ./
RUN go mod download

# Копируем оставшиеся файлы проекта
COPY ../ ./

# Сборка основного приложения
RUN go build -ldflags="-s -w" -o /app/core ./core/main.go

# Финальный этап
FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

# Устанавливаем временную зону
COPY --from=builder /usr/share/zoneinfo/America/New_York /usr/share/zoneinfo/America/New_York

# Копируем собранное приложение
WORKDIR /app
COPY --from=builder /app/core /app/core

# Запускаем приложение
CMD ["/app/core"]
