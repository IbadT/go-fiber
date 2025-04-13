FROM golang:1.23-alpine3.19

WORKDIR /app

# Установка необходимых инструментов
RUN apk add --no-cache git

# Установка CompileDaemon для горячей перезагрузки
RUN go install github.com/githubnemo/CompileDaemon@latest

# Копирование файлов проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка приложения
RUN go build -o main ./cmd/main.go

# Открытие порта
EXPOSE 8000

# Запуск приложения через CompileDaemon для горячей перезагрузки
CMD CompileDaemon --build="go build -o main ./cmd/main.go" --command="./main" 