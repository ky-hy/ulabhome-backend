# ulabhome-backend
全国の大学の研究室プラットフォーム。バックエンドrepo。

# 開発手順

```sh
# イメージビルドおよびコンテナ起動
$ make bu
# コンテナ内に入る
$ make in 
# マイグレーション
$ make migrate
# 初期データ挿入
$ make seed
```

サーバー起動したら、ホスト側で以下のURLでアクセス可能

- API http://localhost:8081
- adminer http://localhost:8082

# swagger

以下のURLで表示可能
[http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

openapi生成ツールは、[https://github.com/swaggo/swag/](https://github.com/swaggo/swag)を利用。

1. 上記公式サイトのドキュメントに則って、コメントを追加し、`$ make swag`を実行。
2. `/docs`ディレクトリにファイルが作成される
3. [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)にアクセス


# フォーマッター、リンター

コンテナ内で以下のコマンドを実行

```sh
# フォーマッター
$ make format
# リンター
$ make lint
```
# テスト

テストを実行するのみであれば以下のコマンドで行う

```sh
$ maek test
```

より詳細なテストのカバレッジを表示したい場合は以下のコマンドを実行

```sh
# coverage.htmlの作成(コンテナ内)
$ make mc
# ブラウザで表示(ホスト側)
$ make wc
```

# DB
開発用データベース情報

|項目|値|
|---|---|
|データベース種類|MySQL|
|サーバ|db|
|ユーザ名|admin|
|パスワード|password|
|データベース|ulabhome|

# マイグレーション

`$ make dry-migrate`でDDLを確認して、`$ make migrate`でマイグレーションする流れとなる

```sh
# マイグレーションする際に発光されるDDLを確認(実行はされない)
$ make dry-migrate
# マイグレーション適用
$ maae migrate
```

# 初期データ挿入

`./_tool/mysql/seed.sql`に初期データを挿入するコマンドを記述

```sh
$ make seed
```
