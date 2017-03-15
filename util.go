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

// LocaleCtx func
func LocaleCtx(ctx context.Context) *Locale {
	return ctx.Value(contextKeyLocale).(*Locale)
}

// ValidateTrans func
func ValidateTrans(locale *Locale, rule, field string) string {
	return locale.T(rule, map[string]interface{}{"Field": locale.T(field)})
}

// ValidateCountTrans func
func ValidateCountTrans(locale *Locale, rule, field string, count int) string {
	return locale.T(rule, map[string]interface{}{"Field": locale.T(field), "Count": count})
}

// ValidateMinMaxTrans func
func ValidateMinMaxTrans(locale *Locale, rule, field string, min, max int) string {
	return locale.T(rule, map[string]interface{}{"Field": locale.T(field), "Min": min, "Max": max})
}
