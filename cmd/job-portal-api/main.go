package main

import (
	"context"
	"fmt"
	"golang/config"
	"golang/internal/auth"
	"golang/internal/cache"
	"golang/internal/database"
	"golang/internal/handlers"
	"golang/internal/repository"
	"golang/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("Welcome to job portal")

}

func StartApp() error {
	cfg := config.GetConfig()

	// privatePem, err := os.ReadFile(`private.pem`)
	// if err != nil {
	// 	return fmt.Errorf("reading private pem %w", err)
	// }
	privatePem := []byte(cfg.AuthConfig.PrivateKey)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		return fmt.Errorf("parsing private key %w", err)
	}

	// publicPem, err := os.ReadFile(`pubkey.pem`)
	// if err != nil {
	// 	return fmt.Errorf("reading public pem %w", err)
	// }
	publicPem := []byte(cfg.AuthConfig.PublicKey)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPem)
	if err != nil {
		return fmt.Errorf("parsing public pem %w", err)
	}

	a, err := auth.NewAuth(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("constructing Auth %w", err)
	}

	log.Info().Msg("main started : initialozing data")
	db, err := database.Open(cfg)
	if err != nil {
		return fmt.Errorf("error in opening the database connection :%w", err)
	}

	pg, err := db.DB()
	if err != nil {
		return fmt.Errorf("error in getting the database instance")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("database is not connected: %w", err)
	}

	rdb := database.ConnectToRedis(cfg)

	redisLayer := cache.NewRDBLayer(rdb)

	// initialize the repository layer
	repo, err := repository.NewRepository(db)
	if err != nil {
		return err
	}

	svc, err := service.NewService(repo, a, redisLayer)
	if err != nil {
		return err
	}

	api := http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.AppConfig.Port),
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handlers.API(a, svc),
	}

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- api.ListenAndServe()
	}()

	shutdownErr := make(chan os.Signal, 1)
	signal.Notify(shutdownErr, os.Interrupt)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server Error %w", err)

	case sig := <-shutdownErr:
		log.Info().Msgf("Main: start Shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := api.Shutdown(ctx) //gracefull shutdown
		if err != nil {
			api.Close() //forceful shutdown
			return fmt.Errorf("server Closed Forcefully %w", err)
		}

	}
	return nil

}
