FROM golang:1.16.3-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache make

WORKDIR /src/challenge
COPY ../.. .
RUN make compile/self-hosted

FROM alpine:3.11

WORKDIR /self_hosted

COPY --from=builder /src/challenge/build/self_hosted /usr/bin/self_hosted

CMD ["self_hosted"]
