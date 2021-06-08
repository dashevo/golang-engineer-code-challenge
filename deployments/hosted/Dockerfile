FROM golang:1.16.3-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache make

WORKDIR /src/challenge
COPY ../.. .
RUN make compile/hosted

FROM alpine:3.11

WORKDIR /hosted

COPY --from=builder /src/challenge/build/hosted /usr/bin/hosted

CMD ["hosted"]
