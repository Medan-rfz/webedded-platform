package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
	auth_repo "webedded.users_management/internal/infrastructure/repositories/auth"
	txexecutor "webedded.users_management/internal/infrastructure/tx_executor"
	auth_service "webedded.users_management/internal/services/auth"
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

	txExecutor, err := txexecutor.NewTxExecutor(conn)
	if err != nil {
		log.Fatalln(err)
	}

	authRepository := auth_repo.NewAuthRepository(conn)
	authService := auth_service.NewUsersService(txExecutor, authRepository)

	// Create user
	authService.Register(ctx, auth_dto.RegisterDTO{
		Email:     "admin@mail.com",
		FirstName: "admin",
		LastName:  "admin",
		Password:  "admin",
	})

	// Set admin role
	// TODO Realize users service
}
