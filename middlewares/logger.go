package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// custom logger middleware
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.StatusCode,
			params.Latency,
		)
	})
}
