FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN chmod +x /app/migrate.sh

CMD ["./migrate.sh","./main"]
