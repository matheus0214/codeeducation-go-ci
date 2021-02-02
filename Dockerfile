FROM golang:1.15-alpine3.12

WORKDIR /app

COPY . .

RUN go build -o math

CMD ["./math"]