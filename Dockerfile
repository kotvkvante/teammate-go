From golang:1.23

WORKDIR /app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
