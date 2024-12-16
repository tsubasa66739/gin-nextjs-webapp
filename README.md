# コンテナ起動

起動と停止

```bash
docker compose --env-file .env up -d
docker compose --env-file .env down

# 起動中のコンテナ確認
docker ps
```

一時停止と再起動

```bash
docker compose stop
docker compose start
```

# DB接続

```bash
psql -h localhost -U postgres
Password for user postgres: ideas
```

# API起動

```bash
cd api
go run main.go
```