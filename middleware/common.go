package middleware

import (
	lgin "github.com/gin-gonic/gin"
	luuid "github.com/google/uuid"
)

func GenerateRequestID() lgin.HandlerFunc {
	return func(pctx *lgin.Context) {
		// Generate a new Request ID
		requestID := "req-" + luuid.New().String()

		// Set Request ID in context
		pctx.Set("requestId", requestID)

		// Add Request ID to response header
		pctx.Writer.Header().Set("X-Request-ID", requestID)

		// Continue processing the request
		pctx.Next()
	}
}
