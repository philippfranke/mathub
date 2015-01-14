FROM golang

RUN apt-get update &&  \
    apt-get install --no-install-recommends -y -q  git && \
    rm -rf /var/lib/apt/lists/*

RUN  git config --global user.email "git@mathub.docker" && git config --global user.name "System Mathub"

ADD . /go/src/github.com/philippfranke/mathub
RUN go get github.com/gorilla/mux && \
    go get github.com/jmoiron/sqlx && \
    go get github.com/go-sql-driver/mysql

ENV MYSQL_USER mathub
ENV MYSQL_PASSWORD mathub
ENV MYSQL_DATABASE mathub

RUN go install github.com/philippfranke/mathub/gateway

RUN apt-get clean

ENTRYPOINT ["/go/bin/gateway", "-data", "/tmp/repos", "-dump", "/go/src/github.com/philippfranke/mathub/datastore/mathub.sql"]
EXPOSE 8080
