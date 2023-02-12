package hlog

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm/logger"
	"time"
)

type GormLoggerConf struct{}

type LoggerForGorm struct{}

func GormLogger(cf *GormLoggerConf) *LoggerForGorm {
	return &LoggerForGorm{}
}

func (l *LoggerForGorm) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *LoggerForGorm) Info(ctx context.Context, s string, i ...interface{}) {
	NewLogger(ctx).Info(s)
}

func (l *LoggerForGorm) Warn(ctx context.Context, s string, i ...interface{}) {
	NewLogger(ctx).Warn(s)
}

func (l *LoggerForGorm) Error(ctx context.Context, s string, i ...interface{}) {
	NewLogger(ctx).Error(errors.New(s), "", i)
}

func (l *LoggerForGorm) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	l.Info(ctx, fmt.Sprintf("[Gorm] %d %s rows: %d\n", elapsed.Milliseconds(), sql, rows))
}
