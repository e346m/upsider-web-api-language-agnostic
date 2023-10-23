package auth

import (
	"context"
	"errors"
	"time"

	"github.com/e346m/upsider-wala/internal/domains"
	"github.com/golang-jwt/jwt/v5"
)

type (
	AuthClient struct {
		secret []byte
	}

	CustomClaims struct {
		JwtData
		*jwt.RegisteredClaims
	}

	JwtData struct {
		Name      string
		BelongsID string
	}
)

const (
	tokenLifeTime        = 1
	refreshTokenLifeTime = 24 * 30
	issuer               = "e346m"
)

func NewAuthClient(secret string) *AuthClient {
	return &AuthClient{
		secret: []byte(secret),
	}
}

func (a *AuthClient) GenerateToken(ctx context.Context, id, name, belongsID string) (string, string, error) {
	now := time.Now()

	accessToken, err := a.generateCustomToken(id, name, belongsID, now, now.Add(time.Hour*tokenLifeTime))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := a.generateCustomToken(id, name, belongsID, now, now.Add(time.Hour*refreshTokenLifeTime))
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (a AuthClient) generateCustomToken(id, name, belongsId string, issuedAt, expiresAt time.Time) (string, error) {
	data := JwtData{
		Name:      name,
		BelongsID: belongsId,
	}

	claims := CustomClaims{
		data,
		&jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   id,
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(a.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthClient) GetPrincipal(
	ctx context.Context,
) (*domains.Principal, error) {
	token, ok := ctx.Value("token").(*jwt.Token)
	if !ok {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims was mapped")
	}
	return &domains.Principal{
		ID:        claims.Subject,
		BelongsID: claims.BelongsID,
	}, nil
}
