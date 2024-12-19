## コンテナ起動

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

### DB接続

```bash
psql -h localhost -U postgres
Password for user postgres: ideas
```

## API

### 起動

```bash
cd api
go run main.go
```

呼び出し例

```bash
# POST
curl -X POST -H 'Content-Type: application/json' -d '{"title":"note test", "body": "hello."}' localhost:8080/api/note

# GET
curl localhost:8080/api/note/1
```

### テスト

モックの生成

```bash
mockgen -source <ディレクトリ>/<対象ファイル>.go -package <ディレクトリ> -destination <ディレクトリ>/<対象ファイル>_mock.go
```

実行

```bash
cd api/

# 全実行
go test ./...

# ディレクトリ配下全て
go test ./service/...
```

### 構成

```bash
api/
├── config            # .envの読み込みなどAPIの設定全般
│   └── config.go
├── controller        # リクエストバリデーション、サービスの呼び出し、レスポンスハンドリング
│   ├── controller.go # ルーターの初期化
│   └── schema        # リクエスト/レスポンスの定義
│       └── schema.go # エラーレスポンスの定義等
├── repository        # DB処理
│   ├── model         # DB定義
│   │   ├── model.go  # モデルの共通項目の定義
│   └── repository.go # DB接続設定
├── service           # 引数（リクエスト等）を元にDB処理を呼び出す
│   └── service.go    # エラーの定義
└── util              # 共通処理
```

呼び出し順

```
controller -> service -> repository
```

utilはどのパッケージからも必要に応じて参照して良い。ただし、utilそのものはどこにも依存してはいけない。
