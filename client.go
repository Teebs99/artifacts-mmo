package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Teebs99/artifacts-mmo/models"
	"io"
	"net/http"
)

type Client struct {
	client *http.Client
	config *models.Config
}

func getClient(config *models.Config) Client {
	return Client{
		client: &http.Client{},
		config: config,
	}
}

func (client *Client) AuthPost(url string, data map[string]string) (map[string]any, error) {
	json_data, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(json_data)))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.config.ApiKey))

	res, err := client.Process(req)

	return res, err
}

func (client *Client) AuthGet(url string) (map[string]any, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.config.ApiKey))

	res, err := client.Process(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (client *Client) Process(req *http.Request) (map[string]any, error) {
	res, err := client.client.Do(req)

	if err != nil && res.StatusCode >= 400 {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return data, err
}
