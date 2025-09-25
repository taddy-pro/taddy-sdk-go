package taddy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type Client struct {
	pubId  string
	apiUrl string
	log    *slog.Logger
}

func New(pubId string, log *slog.Logger) *Client {
	return &Client{pubId: pubId, apiUrl: "https://api.taddy.pro/v1", log: log}
}

func (client *Client) call(method HttpMethod, path string, params any, result any) error {

	client.log.Debug("Request", "method", method, "path", path, "params", params)

	jsonBody, err := json.Marshal(params)
	if err != nil {
		client.log.Error("Error marshalling params", "error", err)
		return err
	}

	var body *bytes.Buffer = nil

	if method == GET {
		path += "?__payload=" + url.QueryEscape(string(jsonBody))
	} else {
		body = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(string(method), client.apiUrl+path, body)
	if err != nil {
		client.log.Error("Failed to create request", "error", err)
		return err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		client.log.Error("Failed to process request", "error", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body2, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		client.log.Error("Failed to read response", "error", err)
		return err
	}

	if resp.StatusCode >= 100 && resp.StatusCode < 300 {
		if resp.StatusCode == 204 {
			client.log.Debug("No content")
			return nil
		}
		err = json.Unmarshal(body2, &result)
		if err != nil {
			client.log.Error("Failed to parse response", "error", err)
			return err
		}
		client.log.Debug("Result", "result", string(body2))
		return nil
	}

	return errors.New(fmt.Sprintf("Error: %d", resp.StatusCode))

}
