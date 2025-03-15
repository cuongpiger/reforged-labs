package middleware

import (
	lgin "github.com/gin-gonic/gin"
	luuid "github.com/google/uuid"
	"net/http"
)

func GenerateRequestID() lgin.HandlerFunc {
	return func(pctx *lgin.Context) {
		// Generate a new Request ID
		requestID := luuid.New().String()

		// Set Request ID in context
		pctx.Set("request_id", requestID)

		// Add Request ID to response header
		pctx.Writer.Header().Set("X-Request-ID", requestID)

		// Continue processing the request
		pctx.Next()
	}
}

// Middleware to check for X-User-ID header
func CheckUserIDHeader() lgin.HandlerFunc {
	return func(pctx *lgin.Context) {
		if pctx.Request.URL.Path == "/healthz" {
			pctx.Next()
			return
		}

		userID := pctx.GetHeader("X-User-ID")
		if userID == "" {
			pctx.JSON(http.StatusBadRequest, lgin.H{"error": "Missing X-User-ID header"})
			pctx.Abort() // Stop further execution
			return
		}

		// Store X-User-ID in the context for later use
		pctx.Set("userID", userID)
		pctx.Next()
	}
}
