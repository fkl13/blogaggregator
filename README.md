# blogaggreagator

A blog aggregator service

## Installation

blogaggreagator need PostgreSQL and Go installed. The database schema is managed with
`goose`. To execute the migration cd into `sql/schema` and run `goose postgres <dsn> up`.

To install the CLI run `go install`.

To configure the program create a `.gatorconfig.json` file in your home directory. 
```json
{"db_url":"postgres://<user>:<password>@<host>/<db-name>","current_user_name":"<username>"}
```

## CLI usage

```console
Commands:
    register    Register new user
    login       Login as user
    reset       Delete all data
    users       List all users
    agg         Fetch feeds in a time interval
    addfeed     Add feed to db
    feeds       List all feeds
    follow      Follow the feed
    following   List the feeds the current users follows
    unfollow    Unfollow a feed
    browse      Browse the latest published feeds
```
