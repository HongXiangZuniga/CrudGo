FROM golang:1.13 AS builder

WORKDIR /app

COPY . .

RUN make install && go build ./cmd/server/api.go

CMD ["./api"]

EXPOSE 8080