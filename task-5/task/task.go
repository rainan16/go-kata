package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// At this point you should be able to create a web server from scratch, or by using `task-4` as a template.
// An interesting use of goroutines could be to have a pool of worker that can run tasks received via a HTTP POST in parallel.
// The following code can get you started to test it out you can use: `curl -X POST -d '{"name":"hello","parallel":2}' 127.0.0.1:8081/task/`
// What about getting something more done in each task?

// Task represent a task to execute in parallel, if you did not try out JSON in
// the previous tasks check out the docs.
// The weird strings after the field declaration are struct tags, and tell the
// json package how to match the json fields with the struct fields.
type Task struct {
	Name     string `json:"name"`
	Parallel int    `json:"parallel"`
}

// Here is a handler that uses some goroutines
func HandleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var (
		task   Task
		result string
		start  = time.Now()
	)
	// read all the body
	buff, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading body: %v", err), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	// ummarshal the body in a variable
	err = json.Unmarshal(buff, &task)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error unmarshaling JSON: %v", err), http.StatusBadRequest)
		return
	}

	if task.Parallel < 1 {
		task.Parallel = 1
	}

	var wg sync.WaitGroup
	wg.Add(task.Parallel)
	// based on how parallel we need to be we run a goroutine
	for i := 0; i < task.Parallel; i++ {
		i := i
		go func() { // func literals are closures
			defer wg.Done()
			fmt.Printf("executing task %q goroutine: %d\n", task.Name, i)
		}()
	}
	// now wait for all the goroutines we started to finish
	wg.Wait()
	result = fmt.Sprintf("task %q execute in %d parallel goroutines\nelapsed time %v\n", task.Name, task.Parallel, time.Since(start))
	fmt.Fprintf(w, "result: %s", result)
}
