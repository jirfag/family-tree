# Family Tree Backend

## 1. Docs

Request|Document
---|---
Restful| <https://www.eolinker.com/#/share/index?shareCode=K2mUw9>
GraphQL| <https://github.com/fredliang44/family-tree/blob/master/schema.graphql>

## 2. URL

Service|URL
---|---
Frontend | <https://fmt.hustunique.com>
Backend | <https://fmt.fredliang.cn>

## 3. Develope & Deploy

1. Edit `config.example.yml` and save it as `config.yml` for develope usage.
2. Then edit `config.example.yml` and save it as `config.deploy.yml` for deploy usage.
3. When deploying, mv `config.deploy.yml` to the same folder with your binary file you've built.

## 4. Project Structure

```shell
tree ./
├── Dockerfile           // docker support
├── README.md
├── config.deploy.yml    // deploy
├── config.example.yml   // example
├── config.yml           // develope config
├── db
│   ├── mongo.go
│   └── redis.go
├── family-tree.go       // main function && router
├── graphql              // handle graphql requests
│   ├── base.go
│   ├── mutation.go
│   ├── query.go
│   ├── resolvers.go
│   ├── schema.go
│   ├── types
│   │   ├── group.go
│   │   ├── project.go
│   │   └── user.go
│   ├── types.go
│   └── utils.go
├── handler              // handler json requests
│   ├── db.go
│   └── register.go
├── middleware
│   ├── auth.go
│   └── cors.go
├── schema.graphql       // graphql doc
└── utils
    ├── config.go        // load config
    └── sms.go           // handle message request with Dayu of Aliyun
```