FROM golang:1.22.5-alpine

WORKDIR /cmd

COPY . .

RUN go mod tidy
RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]