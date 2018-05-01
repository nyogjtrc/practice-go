package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var host = "http://httpbin.org/"

func main() {
	if err := sendGET(); err != nil {
		fmt.Println(err)
	}
}

func sendGET() error {
	fmt.Println("GET", host)
	v := make(url.Values)
	v.Add("id", "123")
	v.Add("enable", "true")

	url := fmt.Sprintf("%s/get?%s", host, v.Encode())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("new request: %s", err)
	}
	req.Header.Set("X-Fake-Header", "hello")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client do: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %s", err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(body))

	return nil
}
