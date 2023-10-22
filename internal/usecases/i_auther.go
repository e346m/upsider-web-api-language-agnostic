package usecases

import (
	"context"
)

type Auther interface {
	GenerateToken(ctx context.Context, id, name, belongsID string) (accessToken, refreshToken string, err error)
}
