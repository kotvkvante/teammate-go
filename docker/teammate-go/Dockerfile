FROM golang:1.24.2

WORKDIR /app

COPY . .
COPY .env .

RUN go build -o main main.go

CMD ["./main"]
