package main

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// Let's build a better logger middleware by using some of the Go most famous
// feature: composition and interfaces.
// Interfaces in Go are a set of methods, a type that implements those methods
// implements the interface, AKA structural typing.
// Composition is used in place of inheritance in Go, we can embed a type in
// another type in Go and, in a way, inherit its methods.
// Back to out middleware, http.ResponseWriter is an interface one of its
// methods is: WriteHeader(code int) what we want to do is capture this call
// and keep the code for later logging, let's see how.

// writerWrapper wraps a http.ResponseWriter and capture the status code
// written to the response.
// http.ResponseWriter is embedded in writerWrapper because is part of the
// struct but it is not given a name. This mean that writerWrapper has all
// the method of http.ResponseWriter.
type writerWrapper struct {
	http.ResponseWriter
	status int
	bytes  int
}

// WriteHeader proxies the call to the underlying http.ResponseWriter and
// capture the status code.
func (w *writerWrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Write proxies the call to the underlying http.ResponseWriter and captures
// the number of bytes written.
func (w *writerWrapper) Write(p []byte) (int, error) {
	n, err := w.ResponseWriter.Write(p)
	w.bytes += n
	return n, err
}

// betterLogs logs each request method, URL and status cod response
func betterLogs(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		wrap := &writerWrapper{
			ResponseWriter: w,
			status:         200,
		}
		next(wrap, r)
		slog.Info(
			"request",
			"code", wrap.status,
			"method", r.Method,
			"url", r.URL.String(),
			"bytes", wrap.bytes,
			"duration", time.Since(started),
		)
	}
}

// 💡Can we log something more?
// Maybe the length of the response?
// Or the time it took for the processing of the request?
// Another method of the http.ResponseWriter is Write([]byte) (int, error).
// Can you proxy this method and capture the length of the response?
// Can you use time.Now() to capture the time before and after the call to next?
// Have a look at the time package https://pkg.go.dev/time!

func measureTime(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		starting := time.Now()
		next(w, r)
		elapsed := time.Since(starting)
		msg := strconv.FormatInt(elapsed.Nanoseconds(), 10)
		slog.Info("timespan", "nanoseconds", msg)
	}
}
