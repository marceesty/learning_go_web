package handlers

import (
	"net/http"
	"github.com/marceesty/learning_go_web/src/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html")
}

