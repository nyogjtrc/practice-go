package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var host = "http://httpbin.org/"

func main() {
	if err := sendGET(); err != nil {
		fmt.Println(err)
	}

	if err := sendPOSTform(); err != nil {
		fmt.Println(err)
	}

	if err := sendPOSTjson(); err != nil {
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

func sendPOSTform() error {
	fmt.Println("POST", host)

	v := make(url.Values)
	v.Add("id", "123")
	v.Add("enable", "true")

	url := fmt.Sprintf("%s/post", host)
	req, err := http.NewRequest("POST", url, strings.NewReader(v.Encode()))
	if err != nil {
		return fmt.Errorf("new request: %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

func sendPOSTjson() error {
	fmt.Println("POST", host)

	jsonString := `{"name":"Timo", "message":"this is a book."}`

	url := fmt.Sprintf("%s/post", host)
	req, err := http.NewRequest("POST", url, strings.NewReader(jsonString))
	if err != nil {
		return fmt.Errorf("new request: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")
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
