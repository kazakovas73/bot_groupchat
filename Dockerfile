FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...
RUN go build -o bot ./cmd/main.go

CMD ["./bot"]
