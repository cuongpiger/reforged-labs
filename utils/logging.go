package utils

import (
	lctx "context"
	lzap "go.uber.org/zap"
)

func GetLogger(pctx lctx.Context) *lzap.Logger {
	if pctx == nil {
		return lzap.L()
	}

	if requestId, ok := pctx.Value("requestId").(string); ok {
		return lzap.L().With(lzap.String("requestId", requestId))
	}

	// Fallback to default logger if not found
	return lzap.L()
}
