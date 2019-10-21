package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	rr := mux.NewRouter()

	rr.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			logger(r.Context()).WithFields(
				logrus.Fields{"status": 200}).Print()
			fmt.Fprint(w, "Hello, World!\n")
		})

	rr.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			logger(r.Context()).WithFields(
				logrus.Fields{"status": 404}).Error()
			w.WriteHeader(404)
		})

	h := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			rr.ServeHTTP(w, r.WithContext(context.WithValue(
				r.Context(), loggerKey,
				logrus.WithFields(logrus.Fields{
					"method": r.Method,
					"url":    r.URL.String(),
					"host":   r.Host,
				}))))
		})

	logrus.Print("Running server on :80")
	logrus.Fatal(http.ListenAndServe(":80", h))
}

type contextKey = string

const loggerKey = contextKey("logger")

func logger(ctx context.Context) *logrus.Entry {
	return ctx.Value(loggerKey).(*logrus.Entry)
}
