package middlewares

import (
	"time"

	"go.uber.org/zap"
)

func LogMiddleware(logger *zap.Logger, timeFormat string, utc bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			path := c.Request.URL.Path

			end := time.Now()
			latency := end.Sub(start)
			if utc {
				end = end.UTC()
			}

			if len(c.Errors) > 0 {
				for _, e := range c.Errors.Errors() {
					logger.Error(e)
				}
			} else {
				logger.Info(path,
					zap.Int("status", c.Writer.Status()),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("time", end.Format(timeFormat)),
					zap.Duration("latency", latency),
				)
			}
			return next(c)
		}
	}
}
