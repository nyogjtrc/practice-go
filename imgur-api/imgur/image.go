package imgur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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
		return nil, parseError(body)
	}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// UploadImage upload a new image
func (ic *Client) UploadImage(path string) (*ResponseImage, error) {
	r := new(ResponseImage)

	req, err := newUploadRequest(fmt.Sprintf("%simage", APIBaseV3), nil, path)
	if err != nil {
		return nil, err
	}

	client, err := ic.HTTPClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = parseResponse(resp, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func newUploadRequest(uri string, params map[string]string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", file.Name())
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, err
}

// ResponseDeleteImage response of delete image result
type ResponseDeleteImage struct {
	Data    bool `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

// DeleteImage deletes an image
func (ic *Client) DeleteImage(id string) (*ResponseDeleteImage, error) {
	client, err := ic.HTTPClient()
	if err != nil {
		return nil, err
	}

	r := new(ResponseDeleteImage)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%simage/%s", APIBaseV3, id), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = parseResponse(resp, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

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

	err = parseResponse(resp, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
