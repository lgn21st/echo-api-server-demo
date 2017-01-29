Echo API Server Demo
====================

System dependencies
------------

- [Golang](https://golang.org)
- [PostgreSQL](https://www.postgresql.org)
- [Glide](https://github.com/Masterminds/glide)
- [Echo](https://github.com/labstack/echo)
- [Gorm](https://github.com/jinzhu/gorm)
- [Sql Migration](https://github.com/rubenv/sql-migrate)


Install
------------

### Create database
The name of database is configurable at `dbconfig.yml`
```shell
createdb echo_demo_dev
```

### Install sql-migration
```shell
go get github.com/rubenv/sql-migrate/...
sql-migrate up
```

### Start Web Server
```shell
go run server.go
```

### Testing through curl
Create user
```shell
curl -X POST \
  http://localhost:8080/users \
  -d 'name=Daniel' \
  -d 'password=MyPassword' \
  -d 'email=daniel@example.com'
```

Deploy to Heroku
------------

```shell
# TODO
```
