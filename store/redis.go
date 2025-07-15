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

func StoreJob(id string) error {
	ctx, cancel := withTimeout()
	defer cancel()
	return rdb.Set(ctx, jobKey(id), utils.StatusQueued, ttl).Err()
}

func UpdateJobStatus(id string, status utils.JobStatus) error {
	ctx, cancel := withTimeout()
	defer cancel()
	return rdb.Set(ctx, jobKey(id), status, ttl).Err()
}

func GetJobStatus(id string) (utils.JobStatus, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	status, err := rdb.Get(ctx, jobKey(id)).Result()
	if err == redis.Nil {
    return "", fmt.Errorf("job not found")
	}
	return utils.JobStatus(status), err
}

func DeleteJob(id string) error {
	ctx, cancel := withTimeout()
	defer cancel()
	return rdb.Del(ctx, jobKey(id)).Err()
}
