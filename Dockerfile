FROM golang:latest

RUN mkdir /app
COPY GoRest /app
WORKDIR /app


CMD [ "/app/GoRest" ]
