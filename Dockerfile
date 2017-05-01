FROM debian:jessie

WORKDIR /app

COPY . /app

ENTRYPOINT [ "./app_linux_amd64" ]
