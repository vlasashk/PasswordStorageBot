FROM golang:1.20-alpine

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./app ./cmd/main/main.go

EXPOSE 8080

CMD ["./app"]