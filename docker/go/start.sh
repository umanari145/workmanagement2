#!/bin/bash

#通常のgo buildなどでは下記のライブラリが入らない
go get github.com/rubenv/sql-migrate/...

cp /go/src/app/dbconfig.yml.sample /go/src/app/dbconfig.yml

sed -i -e "s/__DB_HOST__/$DB_HOST/" /go/src/app/dbconfig.yml
sed -i -e "s/__DB_USER__/$DB_USER/" /go/src/app/dbconfig.yml
sed -i -e "s/__DB_NAME__/$DB_NAME/" /go/src/app/dbconfig.yml
sed -i -e "s/__DB_PASS__/$DB_PASS/" /go/src/app/dbconfig.yml

cd /go/src/app
/go/libs/bin/sql-migrate up