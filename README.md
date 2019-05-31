# Pubbit - A pubbing subbing wonderful friend

Pubbit is a first go at making a PostgreSQL/PubSub streaming update system.

This uses PG's built in triggers and event channels to enable this lightweight go app to publish JSON  update messages to a topic.

## Requirements

1. [Go migrate](https://github.com/golang-migrate/migrate)
1. Go 111 or higher
1. PostgreSQL 9.3 or higher

## Setup

1. Prep the database. This will add an `events` table of `(id: SERIAL, name: VARCHAR)`.
```sh
createdb pubbit
migrate -path ./migrations -database 'postgres://localhost:5432/pubbit?sslmode=disable' up
```

2. Run the listener
```sh
go build
./pubbit run
```

3. Insert data in another prompt
```sh
./pubbit add 'an event'
# 2019/05/31 21:36:03 {"id":1,"name":"an event"}
./pubbit add 'another event'
# 2019/05/31 21:36:39 {"id":10,"name":"another event"}
```
