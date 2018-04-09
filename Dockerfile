FROM ubuntu:16.04

RUN apt-get update \
 && apt-get install -y ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY ./bin/togo-linux-amd64 /app/togo

ENTRYPOINT [ "/app/togo" ]