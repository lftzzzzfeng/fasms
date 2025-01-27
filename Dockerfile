FROM golang:1.23.2-alpine as builder

ENV BUILDER_APP_DIR=/app/srv/fasms

WORKDIR ${BUILDER_APP_DIR}
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o fasms

FROM alpine:3.20.3

ENV APP_DIR=/app/srv/fasms
ENV APP_NAME=fasms

RUN apk update \
  && apk add --no-cache ca-certificates \
  && rm -rf /var/cache/apk/* \
  && adduser -D -h /home/fasms fasms

WORKDIR ${APP_DIR}

COPY --from=builder ${APP_DIR}/${APP_NAME} .
COPY deployment deployment

RUN chown -R fasms:fasms $APP_DIR

USER fasms

CMD ${APP_DIR}/${APP_NAME}
