package users_service

import "context"

type txExecutor interface {
	TxBegin(ctx context.Context, exec func(ctx context.Context) error) error
}

type userService struct {
	txExecutor txExecutor
}

func NewUsersService(txExecutor txExecutor) *userService {
	return &userService{
		txExecutor: txExecutor,
	}
}
