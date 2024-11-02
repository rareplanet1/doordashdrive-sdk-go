package doordashdrive

import (
	"encoding/json"
	"fmt"
	"time"
)

type StoreRequest struct {
	ExternalBusinessID string `json:"external_business_id,omitempty"`
	ExternalStoreID    string `json:"external_store_id,omitempty"`
	Name               string `json:"name,omitempty"`
	PhoneNumber        string `json:"phone_number,omitempty"`
	Address            string `json:"address,omitempty"`
}

func (s *StoreRequest) ToString() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

type StoreResponse struct {
	ExternalBusinessID string     `json:"external_business_id,omitempty"`
	ExternalStoreID    string     `json:"external_store_id,omitempty"`
	Name               string     `json:"name,omitempty"`
	PhoneNumber        string     `json:"phone_number,omitempty"`
	Address            string     `json:"address,omitempty"`
	Status             string     `json:"status,omitempty"`
	IsTest             bool       `json:"is_test,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	LastUpdatedAt      *time.Time `json:"last_updated_at,omitempty"`
}

func (s *StoreResponse) ToString() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

type StoresResponse struct {
	Result            []StoreResponse `json:"result"`
	ContinuationToken string          `json:"continuation_token"`
	ResultCount       int             `json:"result_count"`
}

func (s *StoresResponse) ToString() (string, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (c *Client) GetStore(externalBusinessID string, externalStoreID string) (*StoreResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s/stores/%s", externalBusinessID, externalStoreID)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storeResponse StoreResponse
	err = json.NewDecoder(resp.Body).Decode(&storeResponse)
	if err != nil {
		return nil, err
	}

	return &storeResponse, nil
}

func (c *Client) GetStores(externalBusinessID string) (*StoresResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s/stores", externalBusinessID)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storesResponse StoresResponse
	err = json.NewDecoder(resp.Body).Decode(&storesResponse)
	if err != nil {
		return nil, err
	}

	return &storesResponse, nil
}

func (c *Client) CreateStore(externalBusinessID string, storeRequest StoreRequest) (*StoreResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s/stores", externalBusinessID)
	resp, err := c.doRequest("POST", path, storeRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storeResponse StoreResponse
	err = json.NewDecoder(resp.Body).Decode(&storeResponse)
	if err != nil {
		return nil, err
	}

	return &storeResponse, nil
}

func (c *Client) UpdateStore(externalBusinessID string, externalStoreID string, storeRequest StoreRequest) (*StoreResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s/stores/%s", externalBusinessID, externalStoreID)
	resp, err := c.doRequest("PATCH", path, storeRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storeResponse StoreResponse
	err = json.NewDecoder(resp.Body).Decode(&storeResponse)
	if err != nil {
		return nil, err
	}

	return &storeResponse, nil
}
