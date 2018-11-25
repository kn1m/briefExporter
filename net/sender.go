package net

import (
	"briefExporter/common"
	"io"
	"log"
	"net/http"
)

func logResponse(response *http.Response, body []byte) {
	log.Println("Response status: ", response.Status)
	log.Println("Response headers: ", response.Header)
	log.Println("Response body: ", string(body))
}

func executeRequest(url string, method string, bodyReader io.Reader, headers map[string]string) (*http.Response, error) {
	log.Println("Sending to: ", url)

	req, err := http.NewRequest(method, url, bodyReader)
	common.Check(err)
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	common.Check(err)
	defer resp.Body.Close()

	return resp, err
}
