package users_repo

import (
	queries "webedded.users_management/internal/infrastructure/repositories/users/queries"

	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	conn    *pgx.Conn
	queries *queries.Queries
}

func NewUserRepository(conn *pgx.Conn) *userRepository {
	return &userRepository{
		conn:    conn,
		queries: queries.New(conn),
	}
}
