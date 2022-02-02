module hasura-pg

go 1.17

require (
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	hasura-pg/database v0.0.0-00010101000000-000000000000 // indirect
)

replace hasura-pg/database => ./database
