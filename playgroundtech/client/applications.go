package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetOrder - Returns a specifc order
func (c *Client) GetApplication(applicationID string) (*Application, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/application/%s", c.HostURL, applicationID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Application{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// CreateOrder - Create new order
func (c *Client) CreateApplication(application Application) (*Application, error) {
	rb, err := json.Marshal(application)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/application", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Application{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// UpdateOrder - Updates an order
func (c *Client) UpdateApplication(applicationID string, application Application) (*Application, error) {
	rb, err := json.Marshal(application)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/v1/application/%s", c.HostURL, applicationID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	order := Application{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

// DeleteOrder - Deletes an order
func (c *Client) DeleteApplication(applicationID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/application/%s", c.HostURL, applicationID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
