package utils

import lzap "go.uber.org/zap"

func GoExecute(fn func()) {
	go func() {
		defer recoverPanic()
		fn()
	}()
}

func recoverPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			lzap.L().Error("Recover gorountine panic", lzap.Error(err))
		}
	}
}
