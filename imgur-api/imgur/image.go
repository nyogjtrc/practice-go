package imgur

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Image infomation
type Image struct {
	ID          string
	Title       string
	Description string
	Datetime    int
	Type        string
	Name        string
	Link        string
}

// ResponseImage response image info
type ResponseImage struct {
	Data    Image `json:"data"`
	Success bool  `json:"success"`
	Status  int   `json:"status"`
}

// GetImage get information about an image
func (ic *Client) GetImage(id string) (*ResponseImage, error) {
	client, err := ic.HTTPClient()
	if err != nil {
		return nil, err
	}

	r := new(ResponseImage)

	resp, err := client.Get(fmt.Sprintf("%simage/%s", APIBaseV3, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		return nil, ic.parseError(body)
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

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
		return nil, ic.parseError(body)
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
