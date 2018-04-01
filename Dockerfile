FROM apextoaster/base:master

WORKDIR /app

COPY ./togo /app/togo

ENTRYPOINT [ "/app/togo" ]