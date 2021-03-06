package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!!")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo todo
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the todo title and description only in order to update")
	}

	error := json.Unmarshal(reqBody, &newTodo)

	if error != nil {
		fmt.Fprintf(w, "Invalid data format")
	}

	todos = append(todos, newTodo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func getOneTodo(w http.ResponseWriter, r *http.Request) {
	var todoId = mux.Vars(r)["id"]

	for _, todo := range todos {
		if todoId == todo.ID {
			json.NewEncoder(w).Encode(todo)
		}
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the todo title and description only in order to update")
	}

	todoId := mux.Vars(r)["id"]
	var updateTodo todo

	error := json.Unmarshal(reqBody, &updateTodo)

	if error != nil {
		fmt.Fprintf(w, "Invalid data format")
	}

	for i, singleTodo := range todos {
		if singleTodo.ID == todoId {
			singleTodo.Title = updateTodo.Title
			singleTodo.Description = updateTodo.Description
			todos[i] = singleTodo
			json.NewEncoder(w).Encode(singleTodo)
		}
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["id"]

	for i, todo := range todos {
		if todo.ID == todoId {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Fprintf(w, "Todo with ID %v has been deleted successfully", todoId)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", getOneTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")
	fmt.Println("Starting server at :7070")
	log.Fatal(http.ListenAndServe(":7070", router))
}
