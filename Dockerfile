FROM golang:latest

RUN mkdir /app
COPY GO-REST-Kafka /app
WORKDIR /app


CMD [ "/app/GO-REST-Kafka" ]
