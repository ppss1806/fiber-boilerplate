# Fiber Boilerplate
Simple and scalable boilerplate for Fiber. Structure inspired by [project-layout](https://github.com/golang-standards/project-layout).

## Directory Structure:

```
├── build
│   └── Dockerfile
├── cmd
│   └── example
│       └── main.go
├── config
│   └── example.toml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   └── article.go
│   ├── middlewares
│   │   └── token
│   │       └── token.go
│   ├── models
│   │   └── article_model.go
│   ├── routes
│   │   └── api.go
│   ├── utils
│   │   └── utils.go
│   └── webserver.go
├── LICENSE
├── README.md
└── storage
    ├── private
    │   └── example.html
    ├── private.go
    └── public
```