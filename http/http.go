package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PostJSON .
func PostJSON(url string, js []byte) ([]byte, int, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {

		}
	}()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, nil
}

// PostForm .
func PostForm(url string, form url.Values) ([]byte, int, error) {
	resp, err := http.PostForm(url, form)
	if err != nil {
		return nil, 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {

		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

// Get .
func Get(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {

		}
	}()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
