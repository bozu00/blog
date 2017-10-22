FROM redis:alpine

RUN redis-server --daemonize yes
RUN redis-cli -h localhost sadd bozublog.dev 1

