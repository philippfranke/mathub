mathub
======

![Logo](https://raw.githubusercontent.com/philippfranke/mathub/master/mathubLogo2.jpg?token=AA6Ygt4qooQd1y9V6U9niM8dVfloibyEks5U0D9IwA%3D%3D)

Please use docker for demo or testing!

## Code Status
[![Build Status](https://magnum.travis-ci.com/philippfranke/mathub.svg?token=xJ4sKXa1NvaxvBZ52Ap2&branch=master)](https://magnum.travis-ci.com/philippfranke/mathub)

## Testing/Demo and Deployment
Please use docker in order to review our project. Our docker image only includes our application interface and its dependencies. **Notice: You need your own mysql server, it is not included in our docker image**

The following steps will guide you to run our application:

**1)** mySQL Container 

If you have not installed a mysql server on your machine, please use this docker images. Further information: [here](https://registry.hub.docker.com/_/mysql/)
```
docker pull mysql
docker run --name mathub-mysql -e MYSQL_ROOT_PASSWORD=mathub \
-e MYSQL_USER=mathub -e MYSQL_PASSWORD=mathub -e MYSQL_DATABASE=mathub -d mysql
```

**2)** Build our docker image

Go to the project's root and build a docker image
```
docker build -t mathub:latest .
```

**3)** Run docker image

Run a docker container with a linked mysql container
```
docker run --name mathub -p 80:8080 -v ~/repos:/tmp/repos \ 
           --link mathub-mysql:mysql -d mathub:latest
```
*or*

Run a docker container with your database credentials (replace $your\_mysql\_* with your credentials) 
```
docker run --name mathub -p 80:8080 -v ~/repos:/tmp/repos --link mathub-mysql:mysql \
-e MYSQL_USER=$your_mysql_username -e MYSQL_PASSWORD=$your_mysql_password \
-e MYSQL_DATABASE=$your_mysql_database -e MYSQL_PORT_3306_TCP_ADDR=$your_mysql_ip \
-e MYSQL_PORT_3306_TCP_PORT=$your_mysql_port -d mathub:latest
```

**Reminder: Port Fordwarding on linux: http://localhost, Port Forwarding on others: http://$(\`boot2docker ip\`)**


**4)** Change API host in angular app

Open `/angular/app/scripts/factories/api.js` and change urlBase to "http://localhost" on linux or "http://$(boot2docker_ip)" on windows or mac.

**5)** Run grunt 

Go to `/angular` and run `grunt serve`. Finally start your browser and open `http://localhost:9000`

## Development
For development

### Install dependencies
```
go get github.com/gorilla/mux
go get github.com/jmoiron/sqlx
go get github.com/go-sql-driver/mysql
```

### Change db
Import mathub.sql to your mysql server

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

Default flags:
```
  -listen=:8080           // HTTP listen address
  -data=./repos           // Repository path, all assignments and solutions are under version control
  -dump=""                // mySQL dump file for import ,
```


