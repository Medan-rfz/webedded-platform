package main

import (
	"context"
	"log"
	"os"

	auth_http "webedded.users_management/internal/app/auth/http"
	"webedded.users_management/internal/app/metrics"
	auth_server "webedded.users_management/internal/app/servers/http"
	auth_handler "webedded.users_management/internal/handlers/auth"
	auth_repo "webedded.users_management/internal/infrastructure/repositories/auth"
	txexecutor "webedded.users_management/internal/infrastructure/tx_executor"
	auth_service "webedded.users_management/internal/services/auth"
	users_service "webedded.users_management/internal/services/users"

	"github.com/jackc/pgx/v5"
)

const (
	meterName = "github.com/Medan-rfz/webedded-platform"
	promPort  = 8201
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

	txExecutor, err := txexecutor.NewTxExecutor(conn)
	if err != nil {
		log.Fatalln(err)
	}

	authRepository := auth_repo.NewAuthRepository(conn)
	authService := auth_service.NewUsersService(txExecutor, authRepository)
	usersService := users_service.NewUsersService(txExecutor)
	authHandler := auth_handler.NewAuthHandler(authService, usersService)
	authHttpHandler := auth_http.NewAuthHttpHandler(authHandler)
	authServer := auth_server.NewAuthHttpServer(authHttpHandler)

	go serveMetrics()

	err = authServer.Run(httpConfig)
	if err != nil {
		log.Fatalln(err)
	}
}

func serveMetrics() {
	if err := metrics.Run(promPort); err != nil {
		log.Fatalln("prometheus serve start error")
	}
}
