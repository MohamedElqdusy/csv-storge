package db

import (
	"context"
	"encoding/json"

	"csv-storage/utils"

	"csv-storage/models"

	"github.com/go-redis/redis"
)

type RedisRepository struct {
	conn *redis.Client
}

// Connecting to Redis server
func RedisConnect(url string, password string, database int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       database,
	})
	return client
}

func NewRedis(url string, password string, database int) (*RedisRepository, error) {
	return &RedisRepository{
		RedisConnect(url, password, database),
	}, nil
}

//Creating a Promotion then saving it as JSON to Redis
func (c RedisRepository) CreatePromotion(ctx context.Context, p models.Promotion) error {
	//defer c.Close()

	//Marshal promotion to JSON blob
	blob, err := json.Marshal(p)
	// Save JSON to redis
	err = c.conn.Set("promotions:"+p.Id, blob, 0).Err()
	utils.HandleError(err)
	return nil
}

func (c RedisRepository) FindPromotionById(ctx context.Context, id string) (models.Promotion, error) {
	var promotion models.Promotion
	//defer c.Close()

	reply, err := c.conn.Get("promotions:" + id).Result()
	utils.HandleError(err)
	if err == nil {
		err = json.Unmarshal([]byte(reply), &promotion)
		utils.HandleError(err)
	}
	return promotion, nil
}

func (c RedisRepository) Close() {
	err := c.conn.Close()
	utils.HandleError(err)
}
