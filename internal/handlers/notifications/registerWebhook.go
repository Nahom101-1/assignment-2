package notifications

import (
	"log"
	"net/http"
)

func RegisterWebhook(w http.ResponseWriter, r *http.Request) {
	log.Printf("/POST RegisterWebhook received %s %s\n", r.Method, r.URL.Path)

}
