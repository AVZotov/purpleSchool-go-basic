package api

import (
	"bytes"
	"errors"
	"net/http"
)

const (
	root        = "https://api.jsonbin.io/v3"
	rout        = "/b"
	contentType = "application/json"
)

type Configs interface {
	GetMasterKey() string
}

type LocalStorage interface {
	Create([]byte) error
	Read() error
	Delete(string) error
}

type HttpClient struct {
	client       http.Client
	Key          string
	LocalStorage LocalStorage
}

func NewClient(configs Configs, storage LocalStorage) *HttpClient {
	return &HttpClient{
		Key:          configs.GetMasterKey(),
		LocalStorage: storage,
	}
}

func (client *HttpClient) Create(binName string, body []byte) (*http.Response, error) {
	baseURL := root + rout
	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("X-Bin-Name", binName)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	return resp, nil
}

func (client *HttpClient) Read() {

}

func (client *HttpClient) Update() {

}

func (client *HttpClient) Delete() {

}

func (client *HttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Master-Key", client.Key)
	return client.client.Do(req)
}
