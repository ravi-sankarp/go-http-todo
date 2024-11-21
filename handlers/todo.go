package handlers

import (
	"encoding/json"
	"fmt"
	dbState "http-todo/db"
	"net/http"
	"time"
)

type ResponseObj struct {
	Error bool `json:"error,omitempty"`
	Data  any  `json:"data,omitempty"`
}

type Todo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Due_date    time.Time `json:"due_date"`
	Created_on  time.Time `json:"created_on"`
}

func getTodos(w http.ResponseWriter, req *http.Request) {
	rows, err := dbState.Db.Query("select id, title, description, due_date, created_on from todos")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		response, _ := json.Marshal(ResponseObj{Error: true, Data: err})
		w.Write(response)
		return
	}
	var todos []Todo

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Due_date, &todo.Created_on)
		todos = append(todos, todo)
	}
	w.Header().Add("content-type", "application/json; charset=utf-8")
	response, _ := json.Marshal(ResponseObj{Error: false, Data: todos})
	w.Write(response)
}
func createTodo(w http.ResponseWriter, req *http.Request) {

}
func updateTodo(w http.ResponseWriter, req *http.Request) {

}
func deleteTodo(w http.ResponseWriter, req *http.Request) {

}
