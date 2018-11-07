package ghvisual

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const userAgent = "Sean-Hassett_ghvisual"

type Configuration struct {
	Username string
	Token    string
}

// Makes a HTTP GET request to the passed in URL. Parses JSON data from the body of the response
// and writes matching entries to a passed in struct.
func GetJson(config *Configuration, url string, target interface{}) error {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("User-Agent", userAgent)
	request.SetBasicAuth(config.Username, config.Token)

	response, err := myClient.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	ret, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return json.Unmarshal(ret, &target)
}