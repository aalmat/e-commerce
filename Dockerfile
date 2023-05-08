FROM golang:1.20
WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o e-commerce ./cmd/main.go

EXPOSE 8080

CMD ["./e-commerce"]