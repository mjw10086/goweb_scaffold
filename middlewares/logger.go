package middlewares

import (
	"time"

	"example.com/goweb/mlogger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		mlogger.Logger.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"time":       latency.String(),
		}).Info("Request handled")
	}
}
