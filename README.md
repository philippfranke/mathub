mathub
======

Please fork, instead of working in this repo


## Code Status
[![Build Status](https://magnum.travis-ci.com/philippfranke/mathub.svg?token=xJ4sKXa1NvaxvBZ52Ap2&branch=master)](https://magnum.travis-ci.com/philippfranke/mathub)

## Install
### Install dependencies
```
go get github.com/gorilla/mux
go get github.com/jmoiron/sqlx
go get github.com/go-sql-driver/mysql
```

### Change db
Place update your database credentials in datastore/database.go and import
mathub.sql

### Run or build
__ENV__
```
export MYSQL_USER=mathub
export MYSQL_PASSWORD=mathub
export MYSQL_DATABASE=mathub
export MYSQL_PORT_3306_TCP_ADDR=192.168.59.103
export MYSQL_PORT_3306_TCP_PORT=3306
```

__RUN__
```
cd gateway
go run main.go routes.go
```
__Build__
```
cd gateway
go build
```

Default flags:
```
  -listen=:8080           // HTTP listen address
  -data=./repos           // Repository path
  -dump=""                // mySQL dump for import
```

## Docker Images

0. mySQL Container
```
docker pull mysql
docker run --name mathub-mysql -e MYSQL_ROOT_PASSWORD=mathub -e MYSQL_USER=mathub -e MYSQL_PASSWORD=mathub -e MYSQL_DATABASE=mathub -d mysql
```

1. Build docker image
```
docker build -t mathub:latest .
```

2. Run docker image
```
docker run --name mathub -p 80:8080 --link mathub-mysql:mysql -d mathub:latest
``
