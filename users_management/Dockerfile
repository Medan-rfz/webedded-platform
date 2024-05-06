FROM golang:1.22 as builder
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ./ ./
RUN make build


FROM alpine
COPY --from=builder /go/app /go/
COPY --from=builder /go/bin/ /go/bin/
RUN apk --no-cache add ca-certificates tzdata

CMD ["/go/app"]
EXPOSE 8081
