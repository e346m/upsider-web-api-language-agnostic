package usecases

import (
	"context"

	"github.com/e346m/upsider-wala/internal/domains"
)

func (u *Usecase) SignIn(ctx context.Context, email, password string) (accessToken, refreshToken string, err error) {
	ctx, span := u.tracer.Start(ctx, "SignIn")
	defer span.End()

	// defer checking err to hide hint
	user, _ := u.repo.GetMemberByEmail(ctx, email)
	err = domains.CheckMemberWithPassword(password, user.Password)
	if err != nil {
		span.RecordError(err)
		return
	}

	accessToken, refreshToken, err = u.auther.GenerateToken(ctx, user.ID, user.FullName, user.Organization.ID)

	if err != nil {
		span.RecordError(err)
		return
	}

	return
}
