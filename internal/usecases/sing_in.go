package usecases

import (
	"context"

	"github.com/e346m/upsider-wala/internal/domains"
)

func (u *Usecase) SignIn(ctx context.Context, email, password string) (token string, err error) {
	ctx, span := u.tracer.Start(ctx, "SignIn")
	defer span.End()

	// defer checking err to hide hint
	user, _ := u.repo.GetMemberByEmail(ctx, email)
	err = domains.CheckMemberWithPassword(password, user.Password)
	if err != nil {
		span.RecordError(err)
		return
	}

	// add workflow generating jwt later
	token = "tmp"

	return token, nil
}
