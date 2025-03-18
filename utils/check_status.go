package utils

import (
	"log"
	"net/http"
)

// CheckAPIStatus Sends get request to given url to check status
func CheckAPIStatus(url string) int {
	// Make get request
	res, err := http.Get(url)
	// return http status if service failed
	if err != nil {
		log.Println("Error checking API:", url, err)
		return http.StatusServiceUnavailable
	}
	// close body
	res.Body.Close()
	// return status code if service works/up
	return res.StatusCode
}
