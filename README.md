# Repository Info

## My Bio

- Name : RODERICUS IFO KRISTA
- Stack : Back-End (Golang)

## How to Run

```
$ make start
```

## App Flow

- Register User
- Login User
- Authorized Endpoint:
  - Get List Area
  - Get Team List By Area ID
  - Get Team By ID

## Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/fbab18f85677c9447e6e?action=collection%2Fimport)

## Depedency

- [google-wire](https://github.com/google/wire): for depedency injection
- [jwt](https://github.com/golang-jwt/jwt): for auth process
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt): for hasing password
- [echo](https://echo.labstack.com/): framework
- [gorm](gorm.io/gorm): ORM
- [gorm-postgres](gorm.io/driver/postgres): database driver
