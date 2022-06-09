package featurestores

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//go:generate moq -out httpremote_moq.go . HttpClient
type HttpClient interface {
	Get(string) ([]byte, error)
}

type httpClient struct {
	client *http.Client
}

func newHttpClient() HttpClient {
	return &httpClient{
		client: &http.Client{},
	}
}

func (hc *httpClient) Get(query string) ([]byte, error) {
	get, err := hc.client.Get(query)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

type HttpRemote struct {
	url    string
	client HttpClient
}

func NewRemote(url string) *HttpRemote {
	return &HttpRemote{
		url: url,
	}
}

func (hr *HttpRemote) WithHttpClient(httpClient HttpClient) *HttpRemote {
	hr.client = httpClient
	return hr
}

func (hr *HttpRemote) Build() FeatureStore {
	if hr.client == nil {
		hr.client = newHttpClient()
	}

	return hr
}

type Feature struct {
	Key     string `json:"key"`
	Enabled bool   `json:"enabled"`
}

func (hr *HttpRemote) Get(key string) (bool, error) {
	content, err := hr.client.Get(formatHttpRequestUrl(hr.url, key))
	if err != nil {
		return false, err
	}

	feature, err := unpackFeature(content)
	if err != nil {
		return false, err
	}

	if feature.Key == "" {
		return false, fmt.Errorf("feature was not valid with empty key, aborting")
	}

	return feature.Enabled, nil
}

func unpackFeature(content []byte) (*Feature, error) {
	feature := &Feature{}
	err := json.Unmarshal(content, feature)
	return feature, err
}

func formatHttpRequestUrl(url string, key string) string {
	return fmt.Sprintf("%s/%s", strings.TrimSuffix(url, "/"), key)
}
