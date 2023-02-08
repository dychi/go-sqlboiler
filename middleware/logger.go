package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LogFormatterをカスタマイズ
// https://gin-gonic.com/ja/docs/examples/custom-log-format/
func CustomLogFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage)
	})
}

// Logs a gin HTTP request in JSON format
// https://github.com/sbecker/gin-api-demo/blob/master/middleware/json_logger.go
func JSONLogFormatter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		ctx.Next()

		// Stop timer
		duration := time.Since(start).Milliseconds()

		// JSON
		entry := log.WithFields(log.Fields{
			"duration": duration,
			"method":   ctx.Request.Method,
			"path":     ctx.Request.RequestURI,
			"status":   ctx.Writer.Status(),
			"referrer": ctx.Request.Referer(),
		})

		if ctx.Writer.Status() >= 500 {
			entry.Error(ctx.Errors.String())
		} else {
			entry.Info("")
		}

	}
}
