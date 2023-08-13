FROM alpine:latest

RUN mkdir /app

COPY coreApp /app

CMD [ "/app/coreApp" ]