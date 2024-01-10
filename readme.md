# Token Scraper & API Microservice

This project is composed of two parts: scraper, server
scraper is a cli tool that fetches metadata from ipfs of given cids and stores it into the postgres db.
server is rest api server providing API endpoints for list all tokens and list one token.
This project uses the Ent framework for schema definition and Go's native net/http package and gorilla/mux for routing and handler implementations.
For versioned migration, using golang-migrate

## Prerequisites

- Go 1.14 or newer
- PostgreSQL 15.5 or newer (docker recommended)
- golang-migrate

## Setting up

1. Clone the repo
2. Setup db
3. Run codegen

```
go generate ./ent
```

4. Make .env file
5. Run migration

```
migrate -path migrations -database <DATABASE_URL> up
```

5. Run the scraper

```
go run ./cmd/scraper/main.go scrape -f ipfs_cids.csv
```

6. Run the api server

```
go run ./cmd/server/main.go
```
