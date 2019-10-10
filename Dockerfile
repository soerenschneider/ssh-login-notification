FROM golang:1.13.1-alpine as builder
RUN apk --no-cache add git

COPY . /vgo/
WORKDIR /vgo/

RUN go build -o /sshnot main.go

# Build runtime
FROM alpine:3.10.2 as runtime

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /sshnot /sshnot

# Run as nobody:x:65534:65534:nobody:/:/sbin/nologin
USER 65534

CMD ["/sshnot"]
