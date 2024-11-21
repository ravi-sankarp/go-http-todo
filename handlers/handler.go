package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func TodoHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	switch req.Method {
	case "GET":
		{
			getTodos(w, req)
		}
	case "POST":
		{
			createTodo(w, req)
		}
	case "PUT":
		{
			updateTodo(w, req)
		}
	case "DELETE":
		{
			deleteTodo(w, req)
		}
	default:
		{

			w.WriteHeader(http.StatusMethodNotAllowed)
			res, _ := json.Marshal(ResponseObj{true, "Method not allowed"})
			fmt.Println(ResponseObj{true, "Method not allowed"})
			w.Write(res)
			return
		}
	}

}
