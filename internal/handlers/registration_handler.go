package handlers

import (
	"fmt"
	"net/http"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registration handler")
}
