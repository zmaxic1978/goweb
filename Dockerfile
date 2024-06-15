## указываем, что строим свой образ от образа golang
FROM golang

## указываем рабочую директорию
WORKDIR /usr/src/app

## копируем файл модуля и суммы
COPY go.mod go.sum ./

## Обновляем зависимости пакетов в контейнере. Закачиваем необходимые
RUN go mod tidy

## Копируем файлы проекта в контейнер
COPY . .

## Копируем файлы миграции в контейнер
COPY ./schema ./schema

## далем сборку GO приложения в контейнере
RUN go build ./cmd/main.go

## Установка миграции
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## Резервируем и открываем порт для доступа
EXPOSE 55000

## Выставляем переменные окружения
ENV DB_USERNAME=postgres
ENV DB_PASSWORD=12345

## Запуск миграции и приложения
CMD ["sh", "-c", "migrate -path ./schema -database 'postgres://postgres:12345@db:5432/postgres?sslmode=disable' up && ./main"]
