FROM golang:1.22.3-alpine as builder
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/go-perform-test /app/
WORKDIR /app
EXPOSE 8080
CMD ["./go-perform-test"]