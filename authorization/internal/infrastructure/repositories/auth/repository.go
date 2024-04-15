package auth_repo

import (
	queries "authorization/internal/infrastructure/repositories/auth/queries"

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
