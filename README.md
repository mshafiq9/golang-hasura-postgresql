# golang-hasura-postgresql
GoLang using Hasura migrations and PostgreSQL connectivity

## Installation
Download and install GoLang if you don't have [https://go.dev/doc/install](https://go.dev/doc/install)

#### Prerequisite
This project need an environment variable `DATABASE_URL` for running it successfully. `DATABASE_URL` should follow the patteren `DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}`

```sh
# export DATABASE_URL env var in your shell
# DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}

export DATABASE_URL="<your_posgresql_database_url>"
```

#### Checkout and Run
```sh
$ git clone git@github.com:mshafiq9/golang-hasura-postgresql.git
$ cd golang-hasura-postgresql
$ go run .
```

#### Testing
Note: During testing also cross verify console output when hitting API endpoints.

```sh
# Create Articles:
curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "1", "Title": "title 1", "desc": "description 1", "content": "my articles content 1"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "2", "Title": "title 2", "desc": "description 2", "content": "my articles content 2"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "3", "Title": "title 3", "desc": "description 3", "content": "my articles content 3"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "4", "Title": "title 4", "desc": "description 4", "content": "my articles content 4"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "5", "Title": "title 5", "desc": "description 5", "content": "my articles content 5"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "6", "Title": "title 6", "desc": "description 6", "content": "my articles content 6"}' http://localhost:10000/article

curl -i -X POST -H 'Content-Type: application/json' -d '{"Id": "7", "Title": "title 7", "desc": "description 7", "content": "my articles content 7"}' http://localhost:10000/article
```

```
# Verify Read Articles:
http://localhost:10000/articles
```


## References
https://hasura.io/docs/latest/graphql/core/index.html

https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/

https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
