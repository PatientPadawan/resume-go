package handlers

import (
	"html/template"
	"net/http"
)

type PrintHandler struct{}

func NewPrintHandler() *PrintHandler {
	return &PrintHandler{}
}

func (h *PrintHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/Print-resume.html")
	if err != nil {
		http.Error(w, "Error rendering printable resume", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
