# Riot Go Firebase

RiotとGolangとFirebaseで超簡易的なブログをつくる

## 開発

1. `cp .env.template .env`
2. `npm install`
3. `npm run build`
4. `go run main.go`

## デプロイ

1. `make build`
2. `docker build -t <name> .`

## Docker上のアプリケーションを立ち上げる

`docker run -p 8080:8080 -d <name>`
