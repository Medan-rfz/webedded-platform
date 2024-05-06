package auth_service

import (
	"context"
	"time"

	auth_dto "users_management/internal/domain/dto/auth"
	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"

	"github.com/jackc/pgx/v5"
)

const (
	jwtLiveDuration        = time.Hour * 24
	jwtRefreshLiveDuration = time.Hour * 24 * 30
)

var (
	jwtKey        = []byte("secret-key")
	jwtRefreshKey = []byte("secret-refresh-key")
)

type txExecutor interface {
	TxBegin(ctx context.Context, exec func(ctx context.Context) error) error
}

type authRepo interface {
	CreateUser(ctx context.Context, user auth_entities.User, opts ...repositories.RepoOption) (int64, error)
	GetByEmail(ctx context.Context, email string, opts ...repositories.RepoOption) (*auth_entities.AuthUserModel, error)
	GetAuthByEmail(ctx context.Context, email string, opts ...repositories.RepoOption) (*auth_entities.AuthUserModel, error)
	InsertCredential(ctx context.Context, user auth_entities.UserCredential, opts ...repositories.RepoOption) error
	AddRefreshToken(ctx context.Context, token auth_dto.AddRefreshDTO, opts ...repositories.RepoOption) error
	GetRefreshToken(ctx context.Context, token string, opts ...repositories.RepoOption) (*auth_entities.UserRefreshToken, error)
	NewTx(ctx context.Context) (pgx.Tx, error)
}

type authService struct {
	authRepo   authRepo
	txExecutor txExecutor
}

func NewUsersService(txExecutor txExecutor, authRepo authRepo) *authService {
	return &authService{
		authRepo:   authRepo,
		txExecutor: txExecutor,
	}
}
