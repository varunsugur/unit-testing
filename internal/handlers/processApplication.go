package handlers

import (
	"encoding/json"

	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/models"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ProcessApplication(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("trace id missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("trace id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	id := c.Param("cid")
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error().Msg("give proper cid")
		// c.AbortWithStatusJSON(http.StatusInternalServerError,
		// 	"errors",http.StatusText(http.StatusInternalServerError),
		// })
	}
	var applicationDatas []models.UserApplication

	err = json.NewDecoder(c.Request.Body).Decode(&applicationDatas)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid job details",
		})
	}

	applicationDatas, err = h.service.ProccessApplication(ctx, applicationDatas)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// fmt.Println(jobData)
	c.JSON(http.StatusOK, applicationDatas)

}
