## Recommended VSCode Extension

Go: https://marketplace.visualstudio.com/items?itemName=golang.Go

Prettier: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode

Tailwind CSS IntelliSense: https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss

## 環境設定

`api/.env`

```properties
POSTGRES_VERSION=17
POSTGRES_HOST=localhost
POSTGRES_USER=postgres
POSTGRES_PASS=ideas
POSTGRES_DB=postgres
POSTGRES_PORT=5432
```

`webapp/.env`

```properties
API_URL=http://localhost:8080
```

## コンテナ起動

起動と停止

```bash
docker compose --env-file api/.env up -d
docker compose --env-file api/.env down

# 起動中のコンテナ確認
docker ps
```

一時停止と再起動

```bash
docker compose stop
docker compose start
```

### DB 接続

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

参考に効率化しても

https://zenn.dev/diwamoto/articles/aba45dc2da36b8#%E3%81%93%E3%81%86%E3%81%97%E3%81%9F%E3%82%89%E3%81%A7%E3%81%8D%E3%81%9F

### 構成

```bash
api/
├── cmd               # バッチ処理コマンド
│   ├── cmdOne.go     # バッチ処理本体
│   ├── cmdTwo.go
│   └── root.go       # バッチ処理入口
├── config            # .envの読み込みなどAPIの設定全般
│   └── config.go
├── controller        # リクエストバリデーション、サービスの呼び出し、レスポンスハンドリング
│   ├── controller.go # ルーターの初期化
│   └── schema        # リクエスト/レスポンスの定義
│       └── schema.go # エラーレスポンスの定義等
├── repository        # DB処理
│   ├── model         # DB定義
│   │   └── model.go  # モデルの共通項目の定義
│   └── repository.go # DB接続設定
├── service           # 引数（リクエスト等）を元にDB処理を呼び出す
│   └── service.go    # エラーの定義
├── util              # 共通処理
├── batch_main.go     # バッチのmain
└── main.go           # APIのmain
```

呼び出し順

```
controller -> service -> repository
```

util はどのパッケージからも必要に応じて参照して良い。ただし、util そのものはどこにも依存してはいけない。

## WebApp

パッケージインストール（初回、パッケージ更新時のみ）

```bash
cd webapp
npm i
```

実行

```bash
npm run dev
```

### 構成

```bash
webapp
└── app
    ├── (home)
    │   ├── _component      # / の共通クライアントコンポーネント
    │   │   ├── nav.test.tsx # nav.tsx のテスト
    │   │   └── nav.tsx
    │   ├── (home)
    │   │   ├── home_presentation.tsx  # / のページ本体
    │   │   ├── home_presentation.test.tsx  # home_presentation.tsx のテスト
    │   │   └── page.tsx    # / のサーバ処理
    │   ├── calendar        # /calendar のページ
    │   │   ├── calendar_presentation.tsx  # /calendar のページ本体
    │   │   └── page.tsx    # /calendar のサーバ処理
    │   └── layout.tsx      # / のページレイアウト
    ├── hoge
    │   ├── _component      # /hoge の共通クライアントコンポーネント
    │   ├── (hoge)
    │   │   ├── hoge_presentation.tsx  # /hoge のページ本体
    │   │   ├── hoge_presentation.test.tsx  # hoge_presentation.tsx のテスト
    │   │   └── page.tsx    # /hoge のサーバ処理
    │   ├── fuga            # /hoge/fuga のページ
    │   │   ├── fuga_presentation.tsx  # /hoge/fuga のページ本体
    │   │   ├── fuga_presentation.test.tsx  # fuga_presentation.tsx のテスト
    │   │   └── page.tsx    # /hoge/fuga のサーバ処理
    │   └── layout.tsx      # /hoge のページレイアウト
    ├── _component          # 全体の共通クライアントコンポーネント
    ├── layout.tsx          # 全体のレイアウト
    └── shcema              # リクエスト/レスポンス情報
```
