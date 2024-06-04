package auth_repo

// func (r *authRepository) GetById(ctx context.Context, id int64) (*auth_entities.User, error) {
// 	user, err := r.queries.GetById(ctx, pgtype.Int8{Int64: id, Valid: true})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &auth_entities.User{
// 		Id:           user.UserID.Int64,
// 		Email:        user.Email,
// 		PasswordHash: user.PasswordHash.String,
// 	}, nil
// }
