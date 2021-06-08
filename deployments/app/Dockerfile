FROM golang:1.16.3-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache make

WORKDIR /src/challenge
COPY ../.. .
RUN make compile/app

FROM alpine:3.11

WORKDIR /challenge

ENV TEST_SAMPLE_FILE=/challenge/data.json

COPY --from=builder /src/challenge/build/app /usr/bin/challenge
COPY --from=builder /src/challenge/assets/data.json /challenge

CMD ["challenge"]
