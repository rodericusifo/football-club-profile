# Repository Info

## My Bio

- Name : RODERICUS IFO KRISTA
- Stack : Back-End (Golang)

## Set Up

- Copy <dev|stag|prod|test>.application.yaml.example from folder environments to folder environments with name file like dev.application.yaml and fill the empty value with your environment values

## How to Run

```bash
$ go mod tidy
>  Installing Depedencies...
$ make start-dev
>  Starting Program With Dev Environment...
```

## App Flow

- Register User
- Login User
- Authorized Endpoint:
  - Get List Area
  - Get Team List By Area ID
  - Get Team By ID

## Endpoint Details
| Endpoint                      |               Method                | Info                      | Remark                                 |
| ----------------------------- | :---------------------------------: | :------------------------ | :------------------------------------- |
| /v1/users/register            |               `POST`                | Auth                      | Register User                          |
| /v1/users/login               |               `POST`                | Auth                      | Login User                             |
| /v1/areas                     |               `GET`                 | Areas                     | Get List Areas                         |
| /v1/teams                     |               `GET`                 | Teams                     | Get List Teams                         |
| /v1/teams/:id                 |               `GET`                 | Teams                     | Get Team Detail                        |

## Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/fbab18f85677c9447e6e?action=collection%2Fimport)

## Depedency

- [google-wire](https://github.com/google/wire): for depedency injection
- [jwt](https://github.com/golang-jwt/jwt): for auth process
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt): for hasing password
- [echo](https://echo.labstack.com/): framework
- [gorm](gorm.io/gorm): ORM
- [gorm-postgres](gorm.io/driver/postgres): database driver
