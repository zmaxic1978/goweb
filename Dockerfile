FROM golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 55000

ENV DB_USERNAME=postgres
ENV DB_PASSWORD=12345

CMD ["./main"]