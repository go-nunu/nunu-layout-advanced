# Nunu - 一个优雅的 Golang 脚手架
[英文介绍](https://github.com/go-nunu/nunu-layout-advanced/blob/main/README.md)

Nunu是一个基于Golang的应用脚手架，它的名字来自于英雄联盟中的角色 努努，努努是一个骑在雪怪肩膀上的小男孩，和努努一样，Go-Nunu也是站在巨人的肩膀上，它是由各种第三方库组合而成的，包括gin、gorm、wire、viper、zap、golang-jwt、go-redis、testify、sonyflake、go-survey、cobra等。这些库都是Golang生态中非常流行的库，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。

![Nunu](https://github.com/go-nunu/nunu/blob/main/.github/assets/banner.png)

## 功能

- **Gin**: 一个快速和轻量级的 Golang HTTP web 框架。
- **Gorm**: 一个强大的 Golang ORM 库，支持多种数据库。
- **Wire**: 一个 Golang 编译时依赖注入框架。
- **Viper**: 一个 Golang 配置管理库，支持多种文件格式。
- **Zap**: 一个快速和结构化的 Golang 日志库。
- **Golang-jwt**: 一个 Golang JWT 认证库。
- **Go-redis**: 一个 Golang Redis 客户端库。
- **Testify**: 一个 Golang 测试工具包，提供断言和模拟。
- **Sonyflake**: 一个 Golang 分布式唯一 ID 生成器库。
- **robfig-cron**: 一个很棒的Crontab库。
- More...
## 特性
* **易于使用和定制**：Nunu提供了一个简单直观的API，用于构建Web应用程序。您可以轻松定制应用程序以满足特定需求。
* **高性能和可扩展性**：Nunu旨在具有高性能和可扩展性。它使用最新的技术和最佳实践，确保您的应用程序可以处理高流量和大量数据。
* **安全可靠**：Nunu非常注重安全性。它提供了内置的身份验证、授权和加密支持。它还使用可靠的第三方库，确保您的应用程序安全可靠。
* **模块化和可扩展**：Nunu旨在具有模块化和可扩展性。您可以通过使用第三方库或编写自己的模块轻松添加新功能和功能。
* **文档完善和测试完备**：Nunu文档完善，测试完备。它提供了全面的文档和示例，帮助您快速入门。它还包括一套测试套件，确保您的应用程序按预期工作。
## 要求
要使用Nunu，您需要在系统上安装以下软件：

* Golang 1.16或更高版本
* MySQL5.7或更高版本(可选)
* Redis（可选）
## 安装

你可以通过以下命令来安装Nunu：

```bash
go install github.com/go-nunu/nunu
```

## 使用

使用Nunu非常简单，你只需要按照以下步骤即可：

1. 创建一个新的项目

```bash
nunu new my_project
```

2. 进入项目目录

```bash
cd my_project
```

3. 运行项目

```bash
nunu run
```

4. 在浏览器中打开 http://localhost:8000/ ，你将看到一个欢迎页面。

## 目录结构

Nunu的目录结构如下：

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

## 许可证
Nunu根据MIT许可证获得许可。有关更多信息，请参见LICENSE文件。