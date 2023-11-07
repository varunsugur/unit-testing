package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type key string

const TraceIdKey key = "1"

func (m *Mid) Log() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		traceId := uuid.NewString()

		ctx = context.WithValue(ctx, TraceIdKey, traceId)

		req := c.Request.WithContext(ctx)

		c.Request = req

		c.Next()

	}
}
