
all: build

build:
	go build -o ./app ./cmd


############################################
# Swagger
############################################ 
swagger-gen:
	swag init -g server.go -d "./internal/servers/http,./internal/handlers/auth,./internal/domain/dto/auth"


############################################
# Migrations
############################################ 
migration-up:
	goose -dir migrations postgres ${AUTH_DB_CONN_STR} up

migration-down:
	goose -dir migrations postgres ${AUTH_DB_CONN_STR} down
