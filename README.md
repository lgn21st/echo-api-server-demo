
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
- [x] Store user with encrypted password
- [x] Add unique index to user's email column and avoid duplicated records
- [x] Authentication user
- [ ] Issue JWT token after authentication
- [x] Extract HTTP handler to handler sub-package
- [x] Extract Business logic (e.g. Create User in DB) to service sub-package
- [ ] Update User's password (with JWT token)
- [ ] Create order for user (with JWT token)
- [ ] Introduce a validation framework into application layer
- [ ] Deploy to Heroku
