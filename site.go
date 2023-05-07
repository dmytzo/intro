package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/dmytzo/intro/templates"
)

type Site struct{}

func NewSite() *Site {
	return &Site{}
}

func (s *Site) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpError(w, r, http.StatusMethodNotAllowed, nil)

		return
	}

	if err := renderPage(w, "index"); err != nil {
		httpError(w, r, http.StatusInternalServerError, err)

		return
	}
}

func (s *Site) blogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpError(w, r, http.StatusMethodNotAllowed, nil)

		return
	}

	if err := renderPage(w, "blog"); err != nil {
		httpError(w, r, http.StatusInternalServerError, err)

		return
	}
}

func (s *Site) thoughtsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpError(w, r, http.StatusMethodNotAllowed, nil)

		return
	}

	if err := renderPage(w, "thoughts"); err != nil {
		httpError(w, r, http.StatusInternalServerError, err)

		return
	}
}

func renderPage(w http.ResponseWriter, page string) error {
	tmpl, err := template.ParseFS(templates.FS, "layout.gohtml", fmt.Sprintf("%s.html", page))
	if err != nil {
		return fmt.Errorf("parse files: %w", err)
	}

	if err = tmpl.ExecuteTemplate(w, "layout", map[string]string{"page": page}); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	return nil
}

func httpError(w http.ResponseWriter, r *http.Request, errCode int, err error) {
	statusText := http.StatusText(errCode)

	errMsg := fmt.Sprintf(fmt.Sprintf("[error] %s: %d - %s", r.RequestURI, errCode, statusText))
	if err != nil {
		errMsg = fmt.Sprintf("%s (%s)", errMsg, err.Error())
	}

	log.Printf(errMsg)

	http.Error(w, statusText, errCode)
}
