FROM node:16.13.0 AS FRONT
WORKDIR /web
COPY ./web .
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn install && yarn run build


FROM golang:1.18.4 AS BACK
WORKDIR /go/src/covid
COPY . .
RUN ./build.sh


FROM alpine:latest AS STANDARD

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk add ca-certificates && update-ca-certificates

WORKDIR /
COPY --from=BACK /go/src/covid/server ./server
COPY --from=BACK /go/src/covid/cert/key.pub ./cert/key.pub
COPY --from=FRONT /web/dist ./web/dist
ENTRYPOINT ["/server"]


FROM debian:latest AS db
RUN apt update \
    && apt install -y \
        mariadb-server \
        mariadb-client \
    && rm -rf /var/lib/apt/lists/*


FROM db AS ALLINONE

RUN apt update
RUN apt install -y ca-certificates && update-ca-certificates

WORKDIR /
COPY --from=BACK /go/src/covid/server ./server
COPY --from=BACK /go/src/covid/cert/key.pub ./cert/key.pub
COPY --from=FRONT /web/dist ./web/dist
COPY --from=BACK /go/src/covid/docker-entrypoint.sh /docker-entrypoint.sh

ENTRYPOINT ["/bin/bash"]
CMD ["/docker-entrypoint.sh"]