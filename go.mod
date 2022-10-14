module main

go 1.18

replace api => ./internal/api

replace database => ./internal/database

replace utils => ./utils

require github.com/gorilla/mux v1.8.0

require github.com/lib/pq v1.10.7

require github.com/joho/godotenv v1.4.0
