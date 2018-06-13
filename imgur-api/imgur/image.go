package imgur

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// APIError API error message
type APIError struct {
	Error   string `json:"error"`
	Request string `json:"request"`
	Method  string `json:"method"`
}

// ResponseError reponse message of API error
type ResponseError struct {
	Data    APIError `json:"data"`
	Success bool     `json:"success"`
	Status  int      `json:"status"`
}

// GetImage get information about an image
func (ic *Client) GetImage(id string) {}

// UploadImage upload a new image
func (ic *Client) UploadImage() {}

// DeleteImage deletes an image
func (ic *Client) DeleteImage() {}

// ResponseImageCount response info of my image count
type ResponseImageCount struct {
	Data    int  `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

// MyImageCount get total number of images
func (ic *Client) MyImageCount() (*ResponseImageCount, error) {
	client, err := ic.HTTPClient()
	if err != nil {
		return nil, err
	}

	r := new(ResponseImageCount)

	resp, err := client.Get(fmt.Sprintf("%saccount/me/images/count", APIBaseV3))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		er := new(ResponseError)
		err = json.Unmarshal(body, er)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(er.Data.Error)
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
