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
```

## Todo
- Dockerfile
