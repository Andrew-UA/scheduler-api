package helpers

import (
	"context"
	"errors"
	"scheduler/internal/model"
)

const AuthorizedUser = "authUser"

func GetUserFormContext(ctx context.Context) (model.User, error) {
	authUser, ok := ctx.Value(AuthorizedUser).(model.User)
	if !ok {
		return model.User{}, errors.New("USER NOT EXIST")
	}

	return authUser, nil
}

func SetUserToContext(user model.User, ctx context.Context) context.Context {
	return context.WithValue(ctx, AuthorizedUser, user)
}
