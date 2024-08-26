package httputil

import (
	"context"
	"fmt"

	"github.com/samuelevilla/hasnet-api/internal/types"
)

type userContextKey string

const UserContextKey userContextKey = "user"

func ExtractUserFromContext(ctx context.Context) (*types.User, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}

	user, ok := ctx.Value(UserContextKey).(*types.User)
	if !ok {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

func ContextWithUser(ctx context.Context, user *types.User) context.Context {
	return context.WithValue(ctx, UserContextKey, user)
}
