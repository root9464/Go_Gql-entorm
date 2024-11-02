#!/bin/bash
DOCKERCOMPOSE_PATH="../docker/docker-compose.yml"


start_dbpsq() {
  docker-compose -f $DOCKERCOMPOSE_PATH up -d
}

stop_dbpsq() {
  docker-compose -f $DOCKERCOMPOSE_PATH down
}

rm_dbpsq() {
  docker rm postgres &&
  docker rmi postgres
}

connect_dbpsq() {
  docker exec -it postgres psql -U myuser -d mydb
}

if [ $# -eq 1 ]; then
    "$1"
fi