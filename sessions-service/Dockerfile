FROM golang:1.22 as builder
# RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY ./ ./
RUN make build


FROM alpine
COPY --from=builder /go/bin/ /go/bin/
COPY --from=builder /go/app /go/bin/
COPY --from=builder /go/scripts/ /go/scripts/
COPY --from=builder /go/migrations/ /go/migrations/
COPY --from=builder /go/Makefile /go/Makefile

ENV PATH="/go/bin:${PATH}"
RUN apk --no-cache add make ca-certificates tzdata libc6-compat

ENTRYPOINT [ "/go/scripts/entrypoint.sh" ]

EXPOSE 8081
