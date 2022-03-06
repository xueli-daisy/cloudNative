FROM golang:1.17 AS builder

ENV GOOS=linux
ENV GO111MODULE=off
ENV CGO_ENABLED=0

WORKDIR /build
COPY . .
RUN go build -o HTTPserver simple_server.go

FROM scratch
#COPY --from=builder /build/HTTPserver /httpserver/HTTPserver
COPY --from=builder /build/HTTPserver /
EXPOSE 8080
ENTRYPOINT ["/HTTPserver"]
