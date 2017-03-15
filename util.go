package main

import "context"
import "github.com/jinzhu/gorm"

// RenderCtx func
func RenderCtx(ctx context.Context) *Render {
	return ctx.Value(contextKeyRender).(*Render)
}

// DbCtx func
func DbCtx(ctx context.Context) *gorm.DB {
	return ctx.Value(contextKeyDb).(*gorm.DB)
}
