package handlers

import (
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func API(a *auth.Auth, svc service.UserService) *gin.Engine {
	router := gin.New()

	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Panic().Msg("Middleware not set up")

	}
	h := Handler{
		service: svc,
		auth:    a,
	}

	router.Use(m.Log(), gin.Recovery())

	router.GET("/check", m.Authenticate(check))
	router.POST("/signup", h.Signup)
	router.POST("/login", h.Signin)

	router.POST("/companies", m.Authenticate(h.AddCompany))
	router.GET("/companies", m.Authenticate(h.ViewAllCompanies))
	router.GET("/companies/:id", m.Authenticate(h.ViewCompany))
	router.POST("/companies/job/:cid", m.Authenticate(h.AddJobs))
	router.GET("/companies/jobs/:id", m.Authenticate(h.ViewJob))

	router.GET("jobs", m.Authenticate(h.ViewAllJobs))
	router.GET("/jobs/:id", m.Authenticate(h.ViewJobByID))
	router.POST("/process/application", m.Authenticate(h.ProcessApplication))

	router.POST("/sendotp", h.SendOTP)
	router.POST("/updatepassword", h.UpdatePassword)
	return router
}

func check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Msg": "Good"})

}
