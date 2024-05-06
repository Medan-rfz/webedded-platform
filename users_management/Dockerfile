FROM golang:alpine

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./ ./

RUN make build

CMD ["./app"]
