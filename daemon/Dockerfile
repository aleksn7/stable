FROM golang:latest as builder

WORKDIR /build
COPY ./ /build
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o main ./...

FROM alpine:latest as executor

WORKDIR /usr/bin
COPY --from=builder /build/main /usr/bin/main

ENTRYPOINT [ "/usr/bin/main" ]
