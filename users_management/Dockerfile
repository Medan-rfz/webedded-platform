FROM golang:alpine

RUN apk update && apk add --no-cache make
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./ ./

RUN make build
CMD ["./app"]
