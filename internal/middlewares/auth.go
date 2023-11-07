package middlewares

import (
	"context"
	"errors"
	"golang/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (m *Mid) Authenticate(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		traceID, ok := ctx.Value(TraceIdKey).(string)
		if !ok {
			log.Error().Msg("trace id not present in the context")

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"errors": http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		authHeader := c.Request.Header.Get("Authorization")

		// Splitting the Authorization header based on the space character.
		// Boats "Bearer" and the actual token
		parts := strings.Split(authHeader, " ")
		// Checking the format of the Authorization header
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			log.Error().Err(err).Str("Trace Id", traceID).Send()
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		claims, err := m.a.ValidateToken(parts[1])
		if err != nil {
			log.Error().Err(err).Str("trace id", traceID).Send()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": http.StatusText(http.StatusInternalServerError),
			})
			return
		}

		// If the token is valid, then add it to the context
		ctx = context.WithValue(ctx, auth.Key, claims)

		// Creates a new request with the updated context and assign it back to the gin context
		req := c.Request.WithContext(ctx)
		c.Request = req

		next(c)
	}
}
