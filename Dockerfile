FROM alpine

RUN apk add --update ca-certificates

COPY bin/atenea /usr/bin/atenea

ENV PALERMO_HOST "palermo"
ENV PALERMO_PORT 8003

EXPOSE 3001

ENTRYPOINT atenea -citizens-host=$CITIZENS_HOST -citizens-port=$CITIZENS_PORT -palermo-host=$PALERMO_HOST -palermo-port=$PALERMO_PORT