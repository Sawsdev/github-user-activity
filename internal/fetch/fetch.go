package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
)


func GetFromUrl(url string ) string {

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		fmt.Println("There has been a problem fetching from the url")
	}
	return readResponse(response)
}

func readResponse (response *http.Response) string{
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return ""
	}

	return string(body)
}