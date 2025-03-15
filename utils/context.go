package utils

import (
	lgin "github.com/gin-gonic/gin"
)

// Context wrapper of gin.Context
type Context struct {
	*lgin.Context
}

type HandlerFunc func(*Context)

func WithContext(phandler HandlerFunc) lgin.HandlerFunc {
	return func(ctx *lgin.Context) {
		wrappedContext := &Context{
			ctx,
		}
		phandler(wrappedContext)
	}
}
