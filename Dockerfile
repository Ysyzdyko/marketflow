# Dockerfile

FROM golang:1.21-alpine

WORKDIR /app

# Установим git (если go mod требует)
RUN apk add --no-cache git

# Копируем код и зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Скачиваем wait-for-it
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Собираем приложение
RUN go build -o marketflow ./cmd/app

CMD ["/wait-for-it.sh", "postgres:5432", "--", "/app/marketflow"]
