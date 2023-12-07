package test

import (
	"testing"
	"time"

	redisClient "github.com/saufiroja/redis-go/redis"
	"github.com/stretchr/testify/assert"
)

func TestSetValue(t *testing.T) {
	clientRedis := redisClient.NewRedis()
	err := clientRedis.Set("test_key", "test_value", 10*time.Second)
	assert.NoError(t, err)
}

func TestGetValue(t *testing.T) {
	clientRedis := redisClient.NewRedis()
	expect := "test_value"
	result, err := clientRedis.Get("test_key")
	assert.NoError(t, err)
	assert.Equal(t, expect, result)
}
