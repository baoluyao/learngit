package middleware

import (
	"noticeservice/pkg/app"
	"noticeservice/pkg/errcode"

	"github.com/carr123/fmx"
)

func Recovery() fmx.HandlerFunc {
	return func(c *fmx.Context) {
		defer func() {
			if err := recover(); err != nil {
				app.WriteErrorlog("panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
