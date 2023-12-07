package main

import (
	"net/http"

	"github.com/saufiroja/redis-go/todo"
)

func main() {
	todo := todo.NewTodos()

	http.HandleFunc("/", todo.TodosHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
