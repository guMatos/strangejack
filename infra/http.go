package infra

import (
	"log"
	"net/http"
)

func Get(route string, client http.Client) *http.Response {
	response, err := client.Get(route)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}
