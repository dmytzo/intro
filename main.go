package main

import (
	"log"
	"net/http"

	"github.com/dmytzo/intro/static"
)

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func RequestLogMiddleware(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[request] %s", r.RequestURI)

		handlerFunc(w, r)
	}
}

type ServeMux struct {
	*http.ServeMux

	middlewares []MiddlewareFunc
}

func NewServeMux(middlewares []MiddlewareFunc) *ServeMux {
	return &ServeMux{ServeMux: http.NewServeMux(), middlewares: middlewares}
}

func (m *ServeMux) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerFunc := handler

	for _, middleware := range m.middlewares {
		handlerFunc = middleware(handler)
	}

	m.ServeMux.HandleFunc(pattern, handlerFunc)
}

func (m *ServeMux) HandleFuncRaw(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	m.ServeMux.HandleFunc(pattern, handler)
}

func main() {
	s := NewSite()
	mux := NewServeMux([]MiddlewareFunc{
		RequestLogMiddleware,
	})

	mux.HandleFunc("/", s.indexHandler)
	mux.HandleFunc("/blog", s.blogHandler)
	mux.HandleFunc("/thoughts", s.thoughtsHandler)

	mux.HandleFuncRaw("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-control", "max-age=600")

		http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))).ServeHTTP(w, r)
	})

	log.Println("starting...")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
