package main

import (
	"context"
	"log"
	"os"

	auth_http "users_management/internal/app/auth/http"
	auth_server "users_management/internal/app/servers/http"
	auth_handler "users_management/internal/handlers/auth"
	auth_repo "users_management/internal/infrastructure/repositories/auth"
	auth_service "users_management/internal/services/auth"
	users_service "users_management/internal/services/users"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	connStr, exist := os.LookupEnv("AUTH_DB_CONN_STR")
	if !exist {
		log.Fatalln("DB connection string now found in environment variables")
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	httpConfig := auth_server.AuthHttpServerConfig{
		Address: ":8081",
		IsDev:   true,
	}

	authRepository := auth_repo.NewAuthRepository(conn)
	authService := auth_service.NewUsersService(authRepository)
	usersService := users_service.NewUsersService()
	authHandler := auth_handler.NewAuthHandler(authService, usersService)
	authHttpHandler := auth_http.NewAuthHttpHandler(authHandler)
	authServer := auth_server.NewAuthHttpServer(authHttpHandler)

	err = authServer.Run(httpConfig)
	if err != nil {
		log.Fatalln(err)
	}
}
