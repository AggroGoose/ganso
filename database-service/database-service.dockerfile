FROM alpine:latest

RUN mkdir /app

COPY databaseApp /app

CMD [ "/app/databaseApp" ]