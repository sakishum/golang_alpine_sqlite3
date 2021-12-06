# golang_alpine_sqlite3
I'm trying to build my app which contains go-sqlite3 by docker.

step 1: buil docker image
* docker build -t docker_alpine_exampel .

step 2: run the demo
* docker run -i -t --rm docker_alpine_exampel -v ./data:/go/src/sqlite3_alpine/data:rw
