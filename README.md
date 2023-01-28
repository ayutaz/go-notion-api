# go-notion-api

goでNotionのAPIを使ってみる

# セットアップ

1. clone this repository
2. create .env file for app
3. `docker compose up -d --build`
4. `docker compose exec go-notion-api bash`

# .envファイルの書き方

```dotenv
NOTION_API='notion api key'
NOTION_DATABASE_URL='your notion database url
```

