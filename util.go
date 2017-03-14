package todong

import "context"

// RenderCtx func
func RenderCtx(ctx context.Context) *Render {
	return ctx.Value(contextKeyRender).(*Render)
}
