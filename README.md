
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

### Create Database
The name of database is configurable at `dbconfig.yml`
```shell
createdb echo_demo_dev
```

### Database Schema Migration
```shell
go get github.com/rubenv/sql-migrate/...
sql-migrate up
```

### Start API Server
```shell
go run server.go
```

### How to testing

Install [Postman](https://chrome.google.com/webstore/detail/postman/fhbjgbiflinjbdggehcddcbncdddomop?hl=en) and then import `api.postman_collection`.

### TODO
- [ ] Unit Tests
- [ ] Store user with encrypted password
- [ ] Authentication and issue JWT token
- [ ] Extract HTTP handler to handler sub-package
- [ ] Extract Business (eg. Create User) to service sub-package
- [ ] Update User's password (with JWT token)
- [ ] Create order for user (with JWT token)
- [ ] Deploy to Heroku
