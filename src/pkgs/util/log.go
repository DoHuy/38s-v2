package util

import (
	"go.elastic.co/apm/module/apmzap"
	//"context"
	//"go.elastic.co/apm/module/apmzap"
	//"go.uber.org/zap/zapcore"
	"time"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewProduction(
		zap.WrapCore(
			(&apmzap.Core{}).WrapCore,
		),
	)
	if err != nil {
		panic(err)
	}
	logger = logger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}

// Flush writes out any buffered log entries. Flush should be called before exiting application
func Flush() error {
	return zap.L().Sync()
}

func Zap(logger *zap.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}
		status := c.Writer.Status()

		log := logger.Info

		message := "success"
		if status >= 400 {
			message = "error"
			log = logger.Error
		}

		log(message,
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("time", end.Format(timeFormat)),
			zap.Duration("latency", latency),
		)
	}
}
//
//func Debug(ctx *gin.Context, msg string, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Debug(msg, fields...)
//}
//
//func Info(ctx *gin.Context, msg string, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Info(msg, fields...)
//}
//
//func Error(ctx *gin.Context, msg string, err error, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromError(err, fields...)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Error(msg, fields...)
//}
//
//func Panic(ctx *gin.Context, msg string, err error, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromError(err, fields...)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Panic(msg, fields...)
//}
//
//func Fatal(ctx *gin.Context, msg string, err error, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromError(err, fields...)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Fatal(msg, fields...)
//}
//
//func Warn(ctx common.Context, msg string, fields ...zap.Field) {
//	traceContextFields := buildTraceContextFields(ctx)
//	fields = buildFieldsFromContext(ctx, fields...)
//	zap.L().With(traceContextFields...).Warn(msg, fields...)
//}
//
//func buildTraceContextFields(ctx common.Context) []zapcore.Field {
//	if ctx != nil && ctx.GetGoContext() != nil {
//		return apmzap.TraceContext(ctx.GetGoContext())
//	}
//	return apmzap.TraceContext(context.Background())
//}
//
//func buildFieldsFromContext(ctx common.Context, fields ...zap.Field) []zap.Field {
//	if ctx == nil {
//		return fields
//	}
//
//	if ctx.GetRequestMeta() != nil {
//		fields = append(fields, zap.Object(common.KeyRequestMeta, ctx.GetRequestMeta()))
//	}
//
//	if ctx.GetAppMeta() != nil {
//		fields = append(fields, zap.Object(common.KeyAppMeta, ctx.GetAppMeta()))
//	}
//
//	return fields
//}
//
//func buildFieldsFromError(err error, fields ...zap.Field) []zap.Field {
//	if err == nil {
//		return fields
//	}
//
//	appErr, ok := err.(exception.AppError)
//	errStr := err.Error()
//	if ok {
//		errStr = appErr.Message()
//		fields = append(fields,
//			zap.Int(common.KeyErrorCode, int(appErr.Code())),
//			zap.String(common.KeyErrorType, appErr.ErrorType()))
//	}
//
//	fields = append(fields,
//		zap.String(common.KeyError, errStr),
//	)
//
//	return fields
//}
