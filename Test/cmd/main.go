package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int
	Name        string
	Description string
}

func main() {
	fmt.Println("Server is listening on port 8080")

	tasks := map[int]Task{
		1: Task{
			ID:          1,
			Name:        "Task 1",
			Description: "I need to take a shower",
		},
		2: Task{
			ID:          2,
			Name:        "Task 2",
			Description: "I need to make a homework",
		},
		3: Task{
			ID:          3,
			Name:        "Task 3",
			Description: "I need to take a shower",
		},
		4: Task{
			ID:          4,
			Name:        "Task 4",
			Description: "I need to take a shower",
		},
	}

	router := mux.NewRouter()

	router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		for key, task := range tasks {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
			fmt.Printf("Task ID: %d\n", key)
			fmt.Printf("Name: %s\n", task.Name)
			fmt.Printf("Description: %s\n", task.Description)
			fmt.Println()
		}
	}).Methods("GET")
	router.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskID, ok := vars["id"]
		if !ok {
			http.Error(w, "Missing task ID", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(taskID)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		task, found := tasks[id]
		if !found {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)

	}).Methods("GET")

	router.HandleFunc("/tasl", func(w http.ResponseWriter, r *http.Request) {
		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid JSON request", http.StatusBadRequest)
			return
		}
		newTask.ID = len(tasks) + 1
		tasks[newTask.ID] = newTask
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTask)
	}).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error")
	}

	// Используется такие запросы
	// curl 127.0.0.1:8080/tasks/2
	// curl 127.0.0.1:8080/tasks
	// curl -X POST -H "Content-Type: application/json" -d '{"name": "New Task", "description": "This is a new task"}' 127.0.0.1:8080/tasl

}
