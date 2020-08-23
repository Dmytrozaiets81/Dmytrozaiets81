package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var client *http.Client

func Get(path string) ([]byte, *http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.bilibili.com%s", path), nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, resp, err
}

func init() {
	client = http.DefaultClient
}
