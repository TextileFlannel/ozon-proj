FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /server

ENV PORT=8080

EXPOSE 8080

CMD ["/server"]
