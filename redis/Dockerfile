FROM debian:12.1-slim

EXPOSE 6379

RUN apt update
RUN apt install lsb-release curl gpg -y

RUN curl -fsSL https://packages.redis.io/gpg | gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg

RUN echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/redis.list

RUN apt update
RUN apt install redis -y

ARG REDIS_PASSWORD
ENV REDIS_PASSWORD ${REDIS_PASSWORD}

RUN echo REDIS_PASSWORD

# CMD ["sh", "-c", "redis-server --requirepass $REDIS_PASSWORD"]