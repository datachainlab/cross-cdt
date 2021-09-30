package types

import "context"

type opManagerContextKey struct{}

func OPManagerFromContext(ctx context.Context) OPManager {
	return ctx.Value(opManagerContextKey{}).(OPManager)
}

func ContextWithOPManager(ctx context.Context, opmgr OPManager) context.Context {
	return context.WithValue(ctx, opManagerContextKey{}, opmgr)
}
