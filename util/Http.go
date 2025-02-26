package util

import (
	"io"
	"net/http"
)

func CreateHttpClint(host string, url string) ([]byte, error) {
	ishttp := "http://" + host
	//ishttps := "https://" + host
	client := http.Client{}
	response, err := client.Get(ishttp)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
