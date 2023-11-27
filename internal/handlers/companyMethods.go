package handlers

import (
	"encoding/json"
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/models"
	"strconv"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var comapanyData models.Company

	err := json.NewDecoder(c.Request.Body).Decode(&comapanyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name, location and field",
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(comapanyData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid name, location and field",
		})
		return
	}
	companyData, err := h.service.AddCompanyDetails(ctx, comapanyData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, companyData)

}

func (h *Handler) ViewAllCompanies(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return

	}

	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	companyDetails, err := h.service.ViewAllCompanies(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, companyDetails)

}

func (h *Handler) ViewCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	id := c.Param("id")
	cid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		log.Error().Err(err).Msg("error in parsing cid")
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	companyDetail, err := h.service.ViewCompanyDetail(ctx, cid)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, companyDetail)
}
