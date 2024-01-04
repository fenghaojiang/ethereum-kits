package common

import (
	"context"

	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

func Recoverable(ctx context.Context, fn func(ctx context.Context)) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("recover from panic", zap.Any("err", err), zap.Stack("stack"))
		}
	}()

	fn(ctx)
}
