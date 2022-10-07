FROM golang:1.16-alpine

ADD . /app
WORKDIR /app
RUN go build

CMD ./lifelogspd -config /app/config.json
