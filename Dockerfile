FROM golang:bullseye
WORKDIR /data
COPY /backend .
RUN GOPROXY=https://proxy.golang.com.cn,direct go build -o server


FROM node:bullseye
WORKDIR /data
COPY /web .
RUN npm build

FROM debian:bullseye-slim
WORKDIR /app

COPY --from=0 /data/server server
COPY --from=0 /data/entrypoint.sh entrypoint.sh
COPY --from=0 /data/*.env ./

COPY --from=1 /data/dist www

EXPOSE 80

ENTRYPOINT ["/app/entrypoint.sh"]
