FROM golang:latest

WORKDIR ./

COPY . .

RUN go build -o ./bin/main ./cmd/main.go

EXPOSE 3000

CMD ["./bin/main"]