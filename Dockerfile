FROM golang:alpine3.17

WORKDIR /app

COPY . .

RUN go build

CMD ["./gogoproject"]