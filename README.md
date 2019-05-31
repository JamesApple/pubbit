# Pubbit - A pubbing subbing wonderful friend

Pubbit is a first go at making a PostgreSQL/PubSub streaming update system.

This uses PG's built in triggers and event channels to enable this lightweight go app to publish JSON  update messages to a topic.

## Requirements

1. [Go migrate](https://github.com/golang-migrate/migrate)
1. Go 111 or higher
1. PostgreSQL 9.3 or higher
1. `gcloud` tool with an active service account and topic

## Setup

1. Export your environment variables.
```sh
export GO111MODULE='on'
export PG_URL='postgres://localhost:5432/gopostit?sslmode=disable'
export PROJECT_ID='<YOUR-PROJECT-ID>'
export TOPIC_ID='<TOPIC-NAME-YOU-CREATED>'
```

2. Prep the database. This will add an `events` table of `(id: SERIAL, name: VARCHAR)`.
```sh
createdb pubbit
migrate -path ./migrations -database "$PG_URL" up
```

3. Run the listener
```sh
go build
./pubbit run
```

4. Insert data in another prompt. This will print the result to your console and pubbit to the pubsub.
```sh
./pubbit add 'an event'
# 2019/05/31 21:36:03 {"id":1,"name":"an event"}
# 2019/05/31 21:36:03 Pubbed the data
./pubbit add 'another event'
# 2019/05/31 21:36:39 {"id":10,"name":"another event"}
# 2019/05/31 21:36:39 Pubbed the data
```
