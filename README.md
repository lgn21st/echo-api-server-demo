Echo API Server Demo
====================

Requirements
------------

- [Golang](https://golang.org)
- [PostgreSQL](https://www.postgresql.org)
- [Glide](https://github.com/Masterminds/glide)

Dependencies
------------

- [Echo - Web Framework](https://github.com/labstack/echo)
- [Gorm - ORM for Golang](https://github.com/jinzhu/gorm)

Install
------------

### Create database
The name of database is configurable at dbconfig.yml
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
