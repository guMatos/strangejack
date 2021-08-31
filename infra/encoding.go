package infra

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func DecodeHttpResponse(response http.Response, pointer interface{}) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(body, &pointer)
}
