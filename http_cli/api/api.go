package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"path"
)

const (
	host        = "https://api.jsonbin.io"
	apiVersion  = "v3"
	rout        = "b"
	contentType = "application/json"
)

type Configs interface {
	GetMasterKey() string
}

type HttpClient struct {
	client http.Client
	Key    string
}

func NewClient(configs Configs) *HttpClient {
	return &HttpClient{
		Key: configs.GetMasterKey(),
	}
}

func (client *HttpClient) Create(binName string, body []byte) (*http.Response, error) {
	u, err := getUrl(host, apiVersion, rout)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(body))
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

func (client *HttpClient) Delete(id string) (*http.Response, error) {
	u, err := getUrl(host, apiVersion, rout, id)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodDelete, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *HttpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Master-Key", client.Key)
	return client.client.Do(req)
}

func getUrl(host string, elem ...string) (*url.URL, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(elem...)
	return u, nil
}
