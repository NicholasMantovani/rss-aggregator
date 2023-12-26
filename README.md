# rss-aggregator
This is the sample project from the boot.dev tutorial on youtube

# Dependencies
Thesere are the dev dependency needed for the project

## Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## Install goose
go install github.com/pressly/goose/v3/cmd/goose@latest


# Usage
How to use the external dependencies

## goose
goose postgres postgres://postgres:example@localhost:5432/rssagg up
goose postgres postgres://postgres:example@localhost:5432/rssagg down

## sqlc
sqlc generate