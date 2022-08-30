FROM node:16.13.0 AS web
WORKDIR /web
COPY ./web .
RUN yarn config set registry https://registry.npmmirror.com
RUN yarn install && yarn run build


FROM golang:1.19.0 AS back
WORKDIR /go/src/covid
COPY ./backend .
RUN ./build.sh


FROM alpine:latest AS release

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add curl
RUN apk add ca-certificates && update-ca-certificates

WORKDIR /
COPY --from=back /go/src/covid/server ./server
COPY --from=web /web/dist ./web/dist
ENTRYPOINT ["/server"]
