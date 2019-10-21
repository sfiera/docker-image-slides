package main

import (
	"context"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	rr := mux.NewRouter()

	rr.HandleFunc(
		"/{cow}", func(w http.ResponseWriter, r *http.Request) {
			cmd := exec.Command("cowsay", "-f", mux.Vars(r)["cow"], r.URL.Query().Get("msg"))
			out, err := cmd.CombinedOutput()
			if err != nil {
				logger(r.Context()).WithFields(
					logrus.Fields{"status": 500, "err": err}).Print()
				w.WriteHeader(500)
				return
			}
			logger(r.Context()).WithFields(
				logrus.Fields{"status": 200}).Print()
			w.Write(out)
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
