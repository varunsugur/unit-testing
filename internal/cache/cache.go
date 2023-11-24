package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"golang/internal/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=cache.go -destination=cache_mock.go -package=cache

type RDBLayer struct {
	rdb *redis.Client
}

type Cache interface {
	AddToCache(ctx context.Context, jid uint, jobdata models.Job) error
	GetTheCacheData(ctx context.Context, jid uint) (models.Job, error)

	AddtoOTPCache(ctx context.Context, email string, otp string) error
}

func NewRDBLayer(rdb *redis.Client) Cache {
	return &RDBLayer{
		rdb: rdb,
	}
}

func (r *RDBLayer) AddToCache(ctx context.Context, jid uint, jobdata models.Job) error {
	jobId := strconv.FormatUint(uint64(jid), 10)
	val, err := json.Marshal(jobdata)
	if err != nil {
		return err
	}
	err = r.rdb.Set(ctx, jobId, val, 0).Err()
	return err
}

func (r *RDBLayer) GetTheCacheData(ctx context.Context, jid uint) (models.Job, error) {
	jobId := strconv.FormatUint(uint64(jid), 10)

	val, err := r.rdb.Get(ctx, jobId).Result()
	if err != nil {
		fmt.Println("mmmmmmmmmmmmmmm", err)
		return models.Job{}, err
	}

	var jobData models.Job

	err = json.Unmarshal([]byte(val), &jobData)
	if err != nil {
		fmt.Println("zzzzzzzzzzzzzzzzzzzzzz")
		return models.Job{}, err
	}
	return jobData, nil
}

func (r *RDBLayer) AddtoOTPCache(ctx context.Context, email string, otp string) error {
	err := r.rdb.Set(ctx, email, otp, 5*time.Minute).Err()
	return err
}
