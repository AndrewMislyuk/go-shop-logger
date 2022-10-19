FROM alpine:latest

RUN mkdir /app

COPY grpcApp /app
COPY .env .

CMD [ "/app/grpcApp" ]