# Go Gin Example

## Installation
```
$ go get github.com/DHIMWERAY222/market-place
```

## How to run

### Required

- PostgreSQL
- Redis

### Ready

Create a database called **market_place**

### Conf

You should modify `.env`

```
DB_USER=<username db>
DB_NAME=market_place
DB_PASSWORD=<db password>
DB_HOST=postgres
DB_PORT=<db port>
Xendit_API_KEY=<api key xendit>
RedisAddr=redis:6379

...
```

### Run
```

$ go run app.go
```

Project information and existing API

```
[GIN-debug] GET    /user                     --> market_place.com/controller.GetUser (1 handlers)
[GIN-debug] POST   /user                     --> market_place.com/controller.CreateUser (1 handlers)
[GIN-debug] PUT    /user/:id                 --> market_place.com/controller.UpdateUser (1 handlers)
[GIN-debug] POST   /login                    --> market_place.com/controller.GenerateToken (1 handlers)
[GIN-debug] GET    /product                  --> market_place.com/controller.GetAllProduct (1 handlers)
[GIN-debug] GET    /product/:categoryName    --> market_place.com/controller.GetProductsByCategoryName (1 handlers)
[GIN-debug] POST   /product                  --> market_place.com/controller.CreateProduct (1 handlers)
[GIN-debug] POST   /category                 --> market_place.com/controller.CreateCategory (2 handlers)
[GIN-debug] GET    /category                 --> market_place.com/controller.GetCategory (1 handlers)
[GIN-debug] POST   /cart                     --> market_place.com/controller.AddToCart (2 handlers)
[GIN-debug] GET    /cart                     --> market_place.com/controller.GetAllCartFromUser (2 handlers)
[GIN-debug] DELETE /cart/:id                 --> market_place.com/controller.DeleteCart (2 handlers)
[GIN-debug] POST   /checkout                 --> market_place.com/controller.Payment (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080

```

## Features

- RESTful API
- Gorm
- Jwt-go
- Gin
- App configurable
- Cron
- Redis Cache
