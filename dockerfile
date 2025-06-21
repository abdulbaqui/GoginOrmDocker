FROM golang:1.22 as builder

WORKDIR /app

# Install PostgreSQL client tools
RUN apt-get update && apt-get install -y postgresql-client && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

RUN chmod +x /app/migrate.sh

CMD ["./migrate.sh","./main"]
