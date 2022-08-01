FROM ubuntu:20.04

RUN mkdir /app

WORKDIR /app

COPY ./core/core.exe .

# 下面两个文件，按道理来说是不用 cp 的，但是不 cp 会报错，暂时不知道原因
COPY ./config/app.ini /config/app.ini
COPY ./core/etc/core-api.yaml etc/core-api.yaml
ENTRYPOINT [ "./core.exe","&" ]
EXPOSE 8080