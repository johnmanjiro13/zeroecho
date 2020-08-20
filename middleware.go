package zeroecho

import (
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Config struct {
		Skipper middleware.Skipper
		Logger  *Logger
	}
)

func RequestLogger(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultLoggerConfig.Skipper
	}
	if config.Logger == nil {
		config.Logger = New(os.Stdout, "")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			start := time.Now()
			if err := next(c); err != nil {
				c.Error(err)
			}
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}
			config.Logger.log.Info().
				Str("time", time.Now().Format(time.RFC3339Nano)).
				Str("id", id).
				Str("remote_ip", c.RealIP()).
				Str("host", req.Host).
				Str("uri", req.RequestURI).
				Str("method", req.Method).
				Str("protocol", req.Proto).
				Str("referer", req.Referer()).
				Str("user_agent", req.UserAgent()).
				Int("status", res.Status).
				Str("latency", time.Since(start).String()).
				Send()
			return nil
		}
	}
}
