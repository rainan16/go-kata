package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

// this means all the available IP addresses on port 8080
// if you want your server to bind only one IP you can specify it as IP:PORT
const serverAddress = ":8081"

func main() {
	// we can use the http package default server and we can define our routes
	// by assigning a handler function to a route or path in out URLs
	// each request matching a path will be serve by the corresponding function
	// the matching is longest path first
	// more info https://pkg.go.dev/net/http
	http.HandleFunc("/", measureTime(betterLogs(helloServer)))
	// the {user} element of the path is captured and available in the handler
	http.HandleFunc("/user/{user}/", measureTime(betterLogs(helloUser)))

	// start the webserver
	slog.Info("server running", "address", serverAddress)
	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		slog.Error("ListenAndServe: ", "error", err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	user := r.PathValue("user")
	if user == "coffee" {
		// unfortunately we are permanently a teapot, no coffee for anybody
		w.WriteHeader(http.StatusTeapot)
		fmt.Fprintf(w, "I'm a teapot!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", user)
}

// We used http.HandleFunc to register a handler function to an http route of
// out server. That function signature is:
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)).
// we can have a taste of how we can create some middleware by having a function
// that takes as parameter and returns a func(ResponseWriter, *Request).
// For example here is a middleware that logs the requests.
func logs(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL)
		next(w, r)
	}
}
