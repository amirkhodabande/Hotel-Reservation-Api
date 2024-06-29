package factories

import (
	"context"
	"maps"

	"hotel.com/types"
)

func (f *Factory) CreateUser(data map[string]any) *types.User {
	sample := map[string]any{
		"email":      "test1@gmail.com",
		"first_name": "test1",
		"last_name":  "Ltest1",
		"password":   "password",
	}
	maps.Copy(sample, data)

	p := &types.CreateUserParams{}
	transcode(sample, p)

	user, _ := types.NewUserFromParams(p)

	ctx := context.Background()
	f.UserStore.Insert(ctx, user)

	return user
}
