package handlers

import (
	"encoding/json"
	"errors"
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/models"
	"golang/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type Handler struct {
	service service.UserService
	auth    *auth.Auth
}

func NewHandler(service service.UserService, auth *auth.Auth) (*Handler, error) {
	if service == nil || auth == nil {
		return nil, errors.New("interface and structure cannot be null")
	}
	return &Handler{
		service: service,
		auth:    auth,
	}, nil

}

func (h *Handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing trace id")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return

	}
	var nu models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&nu)

	if err != nil {
		log.Error().Err(err).Str("Trace is ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}
	validate := validator.New()
	err = validate.Struct(nu)

	if err != nil {
		log.Error().Err(err).Msg("could not validate struct")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": "please provide valid username, email and password"})
		return
	}

	userDetails, err := h.service.UserSignup(ctx, nu)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": "please provide valid username, email and password"})
		return
	}
	c.JSON(http.StatusOK, userDetails)
}

func (h *Handler) Signin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing traceid")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	var u models.UserLogin

	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		log.Error().Err(err).Str("Trace is ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}

	token, err := h.service.UserSignin(ctx, u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (h *Handler) SendOTP(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	var data models.VerifyUser

	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		log.Error().Err(err).Str("Trace is ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}

	err = h.service.VerifyUser(ctx, data)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "email sent successfully"})

}

func (h *Handler) UpdatePassword(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	var details models.ResetDetails

	err := json.NewDecoder(c.Request.Body).Decode(&details)
	if err != nil {
		log.Error().Err(err).Str("Trace id ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}

	err = h.service.UpdatePassword(ctx, details)
	if err != nil {
		log.Error().Err(err).Str("traceid", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messege": "password updated successfully"})
}
