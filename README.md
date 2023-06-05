# Nunu - An Elegant Golang Scaffold 

[简体中文介绍](https://github.com/go-nunu/nunu-layout-advanced/blob/main/README_zh.md)

Nunu is an application scaffold based on Golang. Its name comes from the character Nunu in League of Legends, who is a little boy riding on the shoulder of a snowman. Like Nunu, This Project also stands on the shoulders of giants, and it is composed of various third-party libraries, including gin, gorm, wire, viper, zap, golang-jwt, go-redis, testify, sonyflake, go-survey, cobra, etc. These libraries are very popular in the Golang ecosystem, and their combination can help you quickly build an efficient and reliable application.

![Nunu](https://github.com/go-nunu/nunu/blob/main/.github/assets/banner.png)

## Features

- **Gin**: A fast and lightweight Golang HTTP web framework.
- **Gorm**: A powerful Golang ORM library that supports multiple databases.
- **Wire**: A Golang compile-time dependency injection framework.
- **Viper**: A Golang configuration management library that supports multiple file formats.
- **Zap**: A fast and structured Golang logging library.
- **Golang-jwt**: A Golang JWT authentication library.
- **Go-redis**: A Golang Redis client library.
- **Testify**: A Golang testing toolkit that provides assertions and mocking.
- **Sonyflake**: A Golang distributed unique ID generator library.
- **robfig-cron**: A great Crontab library.
- More...
## Features
* **Easy to use and customize**: Nunu provides a simple and intuitive API for building web applications. You can easily customize the application to meet specific needs.
* **High performance and scalability**: Nunu is designed to be high-performance and scalable. It uses the latest technologies and best practices to ensure that your application can handle high traffic and large amounts of data.
* **Secure and reliable**: Nunu places great emphasis on security. It provides built-in authentication, authorization, and encryption support. It also uses reliable third-party libraries to ensure that your application is secure and reliable.
* **Modular and extensible**: Nunu is designed to be modular and extensible. You can easily add new features and functionality by using third-party libraries or writing your own modules.
* **Well-documented and thoroughly tested**: Nunu is well-documented and thoroughly tested. It provides comprehensive documentation and examples to help you get started quickly. It also includes a suite of tests to ensure that your application works as expected.
## Requirements
To use Nunu, you need to install the following software on your system:

* Golang 1.16 or higher
* MySQL 5.7 or higher (optional)
* Redis (optional)
## Installation

You can install Nunu with the following command:

```bash
go install github.com/go-nunu/nunu@latest
```

## Usage

Using Nunu is very simple, you just need to follow these steps:

1. Create a new project

```bash
nunu new my_project
```

2. Enter the project directory

```bash
cd my_project
```

3. Run the project

```bash
nunu run
```

4. Open http://localhost:8000/ in your browser, and you will see a welcome page.

## Directory Structure

The directory structure of Nunu is as follows:


```
.
├── cmd
│   ├── job
│   │   ├── wire
│   │   │   ├── wire.go
│   │   │   └── wire_gen.go
│   │   └── main.go
│   ├── migration
│   │   ├── wire
│   │   │   ├── wire.go
│   │   │   └── wire_gen.go
│   │   └── main.go
│   └── server
│       ├── wire
│       │   ├── wire.go
│       │   └── wire_gen.go
│       └── main.go
├── config
│   ├── local.yml
│   └── prod.yml
├── deploy
│   └── Dockerfile
├── internal
│   ├── database
│   │   └── migration.go
│   ├── handler
│   │   └── user.go
│   ├── job
│   │   └── job.go
│   ├── middleware
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   ├── log.go
│   │   └── sign.go
│   ├── model
│   │   └── user.go
│   ├── provider
│   │   └── provider.go
│   ├── dao
│   │   └── user.go
│   ├── server
│   │   └── http.go
│   └── service
│       ├── user.go
│       └── user_test.go
├── pkg
│   ├── config
│   │   └── config.go
│   ├── db
│   │   └── db.go
│   ├── log
│   │   ├── storage
│   │   │   └── logs
│   │   │       └── server.log
│   │   ├── log.go
│   │   └── log_test.go
│   ├── md5
│   │   └── md5.go
│   ├── rdb
│   │   └── redis.go
│   ├── resp
│   │   └── resp.go
│   ├── sonyflake
│   │   └── sonyflake.go
│   └── uuid
│       └── uuid.go
├── storage
│   └── logs
│       └── server.log
├── test
│   └── server
│       └── handler
│           ├── storage
│           │   └── logs
│           │       └── server.log
│           └── user_test.go
├── web
│   └── index.html
├── LICENSE
├── README.md
├── README_zh.md
├── go.mod
└── go.sum


```

## License
Nunu is licensed under the MIT License. For more information, see the LICENSE file.