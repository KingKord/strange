# base go image
FROM golang:1.21.6-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

# CGO=0 means that we do not use C libraries
RUN CGO_ENABLED=0 go build -o testApp ./cmd/api

# give testApp an executable flag
RUN chmod +x /app/testApp

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/migrations /app
COPY --from=builder /app/testApp /app

CMD [ "/app/testApp" ]