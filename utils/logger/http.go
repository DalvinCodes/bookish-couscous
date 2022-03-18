package logger

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.uber.org/zap"
	"time"
)

type HttpRequest struct {
	Method    string            `json:"method"`
	Headers   map[string]string `json:"headers"`
	Path      string            `json:"path"`
	URL       string            `json:"url"`
	Duration  string            `json:"duration"`
	Payload   []byte            `json:"payload"`
	RequestID string            `json:"requestID"`
	TraceID   string            `json:"traceID"`
	IPAddress string            `json:"ipAddress"`
}

func LoggingMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		if err := c.Next(); err != nil {
			loggingRequest(c, start)
			Panic("Logging Failed", zap.Error(err))
		}
		loggingRequest(c, start)
		return nil
	}
}

func loggingRequest(c *fiber.Ctx, t time.Time) {
	var r HttpRequest

	r.Method = c.Method()
	r.Headers = c.GetReqHeaders()
	r.Path = c.Path()
	r.URL = c.BaseURL()
	r.Duration = time.Since(t).String()
	r.Payload = c.Body()
	r.RequestID = utils.UUIDv4()
	r.IPAddress = c.IP()

	c.Set("X-Request-ID", r.RequestID)

	r.TraceID = c.Get("x-client-trace-id")
	if r.TraceID == "" {
		r.TraceID = utils.UUIDv4()
		c.Set("X_CLIENT_TRACE_ID", r.TraceID)
	}

	Info("Request:", zap.Any("Details", r))

}
