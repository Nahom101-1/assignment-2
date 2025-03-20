package handlers

import (
	"fmt"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("status handler")
}
