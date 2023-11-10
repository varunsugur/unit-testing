package handlers

import (
	"encoding/json"
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/models"
	"net/http"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h *Handler) AddJobs(c *gin.Context) {
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

	id := c.Param("cid")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var jobData models.NewJob

	err = json.NewDecoder(c.Request.Body).Decode(&jobData)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid jov details",
		})
		return
	}
	// validate := validator.New()
	// err = validate.Struct(jobData)

	// if err != nil {
	// 	log.Error().Err(err).Str("trace id", traceId)
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": "please provide valid job details properly"})
	// 	return
	// }

	jobDetails, err := h.service.AddJobDetails(ctx, jobData, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDetails)

}

func (h *Handler) ViewJobByID(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	jid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.ViewJobById(ctx, jid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

func (h *Handler) ViewAllJobs(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	jobDatas, err := h.service.ViewAllJobs(ctx)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobDatas)

}

func (h *Handler) ViewJob(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middlewares.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")

	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	jobData, err := h.service.ViewJob(ctx, cid)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceid)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, jobData)

}

// func (h *Handler) ProcessApplication(c *gin.Context) {
// 	ctx := c.Request.Context()

// 	traceId,ok := ctx.Value(middlewares.TraceIdKey).(string)
// 	if !ok{
// 		log.Error().Msg("trace id missing from context")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
// 			"error": http.StatusText(http.StatusInternalServerError),
// 		})
// 		return
// 	}

// 	_,ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
// 	if !ok{
// 		log.Error().Str("trace id",traceId).Msg("login first")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": http.StatusText(http.StatusInternalServerError),
// 		})
// 		return
// 	}

// 	id := c.Param("cid")
// 	cid,err:= strconv.ParseUint(id,10,64)
// 	if err!=nil{
// 		log.Error().Msg("give proper cid")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
// 			"errors":http.StatusText(http.StatusInternalServerError),
// 		})
// 	}
// 	var jobData models.Job

// 	err =json.NewDecoder(c.Request.Body).Decode(&jobData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceId)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "please provide valid job details",
// 		})
// 		return
// 	}

// 	h.service.

// }
