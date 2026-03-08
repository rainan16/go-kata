package main

import (
	"fmt"
	"net/http"
	"task-5/task"
)

func main() {
	// Register the task handler
	http.HandleFunc("/task/", task.HandleTask)

	// Start the server
	fmt.Println("Server starting on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
