# workmanagement2

サーバーサイド go
フロントエンド vue


### vueに関して

開発時(ホットリーロード)
```
npm run serve
```

ビルド
```
npm run build
```

構文チェック
```
npm run lint
```

### ルーティング

#### 勤怠報告がめん
http://localhost:8080/home

#### 予約画面
http://localhost:8080/reserve



### APIルーティング
gorilla/mux
https://qiita.com/stranger1989/items/7d95778d26d34fd1ddef


### マイグレーションツール
sql-migrate
https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0

コマンド
```
#help
/go/libs/bin/sql-migrate --help

#確認
/go/libs/bin/sql-migrate status

#マイグレーションファイル作成
go/libs/bin/sql-migrate new ファイル名

#マイグレーション実行
go/libs/bin/sql-migrate up

#マイグレーション戻し(1stepのみ)
go/libs/bin/sql-migrate down

```

### その他のメモ
@vue/cli 4.5.3で動作確認

sql-migrationのインストールに時間がかかるため、docker-compose up (-dなしの方がわかりやすい)で30秒ぐらい立って
マイグレーションが実行された後で実行すること