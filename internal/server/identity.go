package server

import "context"

type contextKey string

const userIDKey contextKey = "user-id"

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func UserIDFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(userIDKey)
	if v == nil {
		return "", false
	}
	userID, ok := v.(string)
	return userID, ok
}
