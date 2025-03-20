package handlers

import (
	"net/http"
)

func LocalHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
