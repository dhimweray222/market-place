# Market Place

## Installation
```
$ go get github.com/dhimweray222/market-place
```
## ERD
```
![alt text](https://github.com/dhimweray222/market-place/blob/main/public/market%20place.png)  
```

## Structure
```
├── config
│   └──cache.go // redis config
|	  └── db.go // db config
|	  └── registry.go // table registry
├── controller
│   └── auth.go // session Controller
│   └── cart.go // shopping cart Controller
│   └── category.go // category Controller
│   └── checkout.go // checkout Controller
│   └── payment.go // payment Controller
│   └── product.go // product Controller
│   └── user.go // user Controller
├── models
│   └── user.go // user model
│   └── category.go // category model
│   └── product.go // product model
│   └── cart.go // cart model
├── Routers
│   └── auth.go // session Routes
│   └── cart.go // shopping cart Routes
│   └── category.go // category Routes
│   └── checkout.go // checkout and payment Routes
│   └── product.go // product Routes
│   └── user.go // user Routes
└── app.go
```

## How to run

### Required

- PostgreSQL
- Redis

### Ready

Create a database called **market_place**

### Conf

You should modify `.env`

```![Alt text](image link)
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

## API

#### /user
* `POST` : Create a new User

#### /login
* `POST` : login  User

#### /product
* `POST` : Create a product
* `GET` : Get all products

#### /product/:categoryName
* `GET` : Get all products by category

#### /category
* `POST` : Create a category
* `GET` : Get all categories

#### /cart
* `POST` : Add product to cart
* `GET` : get all carts from logged in user

#### /cart/:id
* `DELETE` : Delete cart from logged in user by product id

#### /checkout
* `POST` : Checkout and make a payment from logged in user
