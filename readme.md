//Create migration

migrate create -ext sql -dir migrations create_table_users

//Migrate table
migrate -database "postgres://user:password@host:port/database_name?sslmode=disable" -path migrations up