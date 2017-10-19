
```
goose up 
```

[api-app](http://localhost:1323/)

[redis-commander](http://localhost:8080/)

[swagger-ui](http://localhost:8080/)

[swagger-editor](http://localhost:8080/)

```
removecontainers() {
    docker stop $(docker ps -aq)
    docker rm $(docker ps -aq)
}

armaggedon() {
    removecontainers
    docker network prune -f
    docker rmi -f $(docker images --filter dangling=true -qa)
    docker volume rm $(docker volume ls --filter dangling=true -q)
    docker rmi -f $(docker images -qa)
}
```
