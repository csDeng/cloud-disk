FROM redis:5.0

LABEL maintainer="github.com/csDeng"

WORKDIR /data/redis

COPY ./redis.conf /etc/redis/redis.conf

CMD ["redis-server", "/etc/redis/redis.conf"]
EXPOSE 6379

