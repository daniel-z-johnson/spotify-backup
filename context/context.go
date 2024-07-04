package context

import "context"

const (
	sessionKey = "session"
	codeKey    = "code"
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

func WithCode(ctx context.Context, code *string) context.Context {
	return context.WithValue(ctx, codeKey, code)
}

func GetCode(ctx context.Context) *string {
	val := ctx.Value(codeKey)
	code, ok := val.(*string)
	if !ok {
		return nil
	}
	return code
}
