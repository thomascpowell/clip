package store

import (
	"fmt"
	"context"
	"video-api/utils"
	"time"
	"github.com/redis/go-redis/v9"
)

const ttl = time.Hour
const redisTimeout = 3 * time.Second
const jobKeyPrefix = "job:"
var rdb *redis.Client

func withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), redisTimeout)
}

func jobKey(id string) string {
	return jobKeyPrefix+id
}

func hsetWithTTL(key string, values map[string]any) error {
	ctx, cancel := withTimeout()
	defer cancel()
	if err := rdb.HSet(ctx, key, values).Err(); err != nil {
		return err
	}
	return rdb.Expire(ctx, key, ttl).Err()
}

// start redis on a provided address
func InitRedis(addr string) error {
	ctx, cancel := withTimeout()
	defer cancel()
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := rdb.Ping(ctx).Result()
	return err
}

func StoreJob(id, format string) error {
	return hsetWithTTL(jobKey(id), map[string]any{
		"status": utils.StatusQueued,
		"format": format,
	})
}

func UpdateJobStatus(id string, status utils.JobStatus) error {
	return hsetWithTTL(jobKey(id), map[string]any{
		"status": status,
	})
}

func DeleteJob(id string) error {
	ctx, cancel := withTimeout()
	defer cancel()
	return rdb.Del(ctx, jobKey(id)).Err()
}

func GetStatusAndFormat(id string) (utils.JobStatus, string, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	result, err := rdb.HMGet(ctx, jobKey(id), "status", "format").Result()
	if err != nil {
		return "", "", err
	}
	status := result[0]
  format := result[1]
	if status == nil || format == nil {
		return "", "", fmt.Errorf("missing job fields")
	}
	return utils.JobStatus(status.(string)), format.(string), nil
}

