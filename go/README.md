# monorepo/server

![CI/CD](https://github.com/ispec-inc/monorepo/workflows/gotest/badge.svg)
[![codecov](https://codecov.io/gh/ispec-inc/monorepo/branch/master/graph/badge.svg)](https://codecov.io/gh/ispec-inc/monorepo)

単一のDBに対してやりとりをするようなシステムアーキテクチャにおけるGoのサーバーのテンプレートPJです．

## Usage
1. 以下のコマンドを実行
```
$ make build
$ make init
$ make article-migrate in ../migration
$ make seed
$ make server
```

2. ヘルスチェック
```
$ curl localhost:9000/health
{"message":"success"}
```

## Architecture
Modular Monolithicのアーキテクチャを採用している。
```
├── cmd
│   ├── db
│   │   └── seed
│   │       └── main.go
│   └── server
│       └── main.go # エントリーポイント
├── docker
│   ├── dev
│   │   ├── Dockerfile # 開発用のDockerfile
│   │   └── my.cnf
│   └── prod
│       └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg # svcやcmdから参照される便利な共通pkg群
├── proto # protocで生成されたコード群
└── svc # 互いに疎結合なサービス群 pkg/routerでそれぞれのサービスがアグリゲートされる
    ├── admin
    ├── article
    ├── media
    └── notification
```
