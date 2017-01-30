
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

Step 1. Create User Account
```shell
curl --request POST \
  --url http://localhost:8080/users \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data 'name=Daniel&email=daniel%40example.com&password=MyPassword'
```

Step 2. Auth and retrieve JWT Token
```shell
curl --request POST \
  --url http://localhost:8080/auth \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data 'email=daniel%40example.com&password=MyPassword'
```

Step 3. Update User's name with JWT token
```shell
curl --request PUT \
  --url http://localhost:8080/update_name \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhbmllbEBleGFtcGxlLmNvbSJ9.WGtSoCzB6TAOIJED5FWuvGbE_wiI9UVBv4BGHXvW_Og' \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data name=Daniel%20Lv
```


### TODO
- [ ] Unit Tests
- [x] Store user with encrypted password
- [x] Add unique index to user's email column and avoid duplicated records
- [x] Authentication user
- [x] Issue JWT token after authentication
- [x] Extract HTTP handler to handler sub-package
- [x] Extract Business logic (e.g. Create User in DB) to service sub-package
- [x] Update User's name (with JWT token)
- [ ] Create order for user (with JWT token)
- [ ] Introduce a validation framework into application layer
- [ ] Deploy to Heroku
