package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	redisClient "github.com/saufiroja/redis-go/redis"
)

type ITodos interface {
	TodosHandler(w http.ResponseWriter, r *http.Request)
}

type Todos struct {
	redis redisClient.IRedis
}

func NewTodos() ITodos {
	return &Todos{
		redis: redisClient.NewRedis(),
	}
}

func (t *Todos) TodosHandler(w http.ResponseWriter, r *http.Request) {
	data := Todo{
		ID:          1,
		Name:        "john",
		Description: "do something",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("error marshaling todo data to JSON:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = t.redis.Get("todo")
	if err == redis.Nil {
		log.Println("todo data not found in Redis, setting it now")
		err := t.redis.Set("todo", string(jsonData), 3*time.Minute)
		if err != nil {
			log.Println("error setting todo data in Redis:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	log.Println("todo data found in Redis, returning it now")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("error writing JSON response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Todo represents the todo data structure
type Todo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
