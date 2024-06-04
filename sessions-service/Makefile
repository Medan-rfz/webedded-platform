
all: build

build:
	go build -0 ./db_init ./cmd/db_init
	go build -o ./app ./cmd/app


############################################
# Sqlc
############################################ 
sqlc-gen:
	cd ./sqlc && sqlc generate

############################################
# Swagger
############################################ 
swagger-gen:
	swag init -g server.go -d "./internal/app/servers/http,./internal/app/auth/http,./internal/domain/dto/auth"


############################################
# Migrations
############################################ 
migration-up:
	goose -dir migrations postgres ${AUTH_DB_CONN_STR} up

migration-down:
	goose -dir migrations postgres ${AUTH_DB_CONN_STR} down
