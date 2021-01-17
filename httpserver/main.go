package main

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "gopher academy smaple")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logrus.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

