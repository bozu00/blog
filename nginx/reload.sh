#!/bin/bash

docker exec -it  blog_nginx_1  /usr/local/openresty/nginx/sbin/nginx -s reload
