package main

import (
	"fmt"
	"http-request/pekostruct"
	"io/ioutil"
	"log"
	"net/http"
)

func bodyParser(resp *http.Response) []byte {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body

}

func main() {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")

	if err != nil {
		log.Fatal(err)
	}

	pekostruct, err := pekostruct.UnmarshalPekostruct(bodyParser(resp))
	fmt.Println(pekostruct.BaseExperience)

	// fmt.Println(bodyParser(resp))
}
