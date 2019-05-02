package views

import (
	"log"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request) error

func HE(hdlr HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := hdlr(w, req); err != nil {
			switch e := err.(type) {
			case IStatusError:
				// We can retrieve the status here and write out a specific
				// HTTP status code.
				if e.Error != nil {
					log.Printf("HTTP %d - %s", e.Status(), e)
					errview := NewView("main", "err.tmpl")
					errview.Vars["Error"] = e.Error()
					w.WriteHeader(e.Status())
					errview.Render(w)
				}
			default:
				// Any error types we don't specifically look out for default
				// to serving a HTTP 500
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
			return
		}
	}
}
