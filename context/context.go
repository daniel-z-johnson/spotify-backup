package context

import "context"

const (
	sessionKey = "session"
)

func WithSession(ctx context.Context, code *string) context.Context {
	return context.WithValue(ctx, sessionKey, code)
}

func GetSession(ctx context.Context) *string {
	val := ctx.Value(sessionKey)
	code, ok := val.(*string)
	if !ok {
		return nil
	}
	return code
}
