package main

import (
	"fmt"
	logrus "github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	//http.Handle("/go",gopherMiddleware(http.HandlerFunc(gopher)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "package main #14")
}

//func gopher(w http.ResponseWriter, req *http.Request)  {
//	fmt.Fprintf(w,"hi gophery")
//}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logrus.Infof("uri: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}

//func gopherMiddleware(next http.Handler) http.Handler  {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		logrus.Infof("gopher: %s",r.RequestURI)
//		next.ServeHTTP(w,r)
//	})
//}