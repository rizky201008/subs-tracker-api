FROM golang:1.23-alpine
LABEL authors="vixiloc"

WORKDIR /app

ENV APP_PORT=3000

COPY . .

RUN go get

RUN go build -o apps

CMD "./apps"

EXPOSE ${APP_PORT}