package doordashdrive

import (
	"encoding/json"
	"fmt"
	"time"
)

type BusinessRequest struct {
	ExternalBusinessID string `json:"external_business_id,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	ActivationStatus   string `json:"activation_status,omitempty"`
}

type BusinessResponse struct {
	ExternalBusinessID string                   `json:"external_business_id,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Description        string                   `json:"description,omitempty"`
	ActivationStatus   string                   `json:"activation_status,omitempty"`
	CreatedAt          *time.Time               `json:"created_at,omitempty"`
	LastUpdatedAt      *time.Time               `json:"last_updated_at,omitempty"`
	IsTest             bool                     `json:"is_test,omitempty"`
	ExternalMetadata   BusinessExternalMetadata `json:"external_metadata,omitempty"`
}

type BusinessesResponse struct {
	Result            []BusinessResponse `json:"result"`
	ContinuationToken string             `json:"continuation_token"`
	ResultCount       int                `json:"result_count"`
}

type BusinessExternalMetadata struct {
	NumberOfStores    int      `json:"number_of_stores"`
	ClientEmail       string   `json:"client_email,omitempty"`
	ClientPhoneNumber string   `json:"client_phone_number,omitempty"`
	ExternalStoreIds  []string `json:"external_store_ids,omitempty"`
}

func (c *Client) GetBusiness(externalBusinessID string) (*BusinessResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s", externalBusinessID)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var businessResponse BusinessResponse
	err = json.NewDecoder(resp.Body).Decode(&businessResponse)
	if err != nil {
		return nil, err
	}

	return &businessResponse, nil
}

func (c *Client) GetBusinesses() (*BusinessesResponse, error) {
	resp, err := c.doRequest("GET", "/developer/v1/businesses", nil)
	if err != nil {
		return nil, err
	}

	var businessesResponse BusinessesResponse
	err = json.NewDecoder(resp.Body).Decode(&businessesResponse)
	if err != nil {
		return nil, err
	}

	return &businessesResponse, nil
}

func (c *Client) CreateBusiness(businessRequest BusinessRequest) (*BusinessResponse, error) {
	resp, err := c.doRequest("POST", "/developer/v1/businesses", businessRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var businessResponse BusinessResponse
	err = json.NewDecoder(resp.Body).Decode(&businessResponse)
	if err != nil {
		return nil, err
	}

	return &businessResponse, nil
}

func (c *Client) UpdateBusiness(externalBusinessID string, businessRequest BusinessRequest) (*BusinessResponse, error) {
	path := fmt.Sprintf("/developer/v1/businesses/%s", externalBusinessID)
	resp, err := c.doRequest("PATCH", path, businessRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var businessResponse BusinessResponse
	err = json.NewDecoder(resp.Body).Decode(&businessResponse)
	if err != nil {
		return nil, err
	}

	return &businessResponse, nil
}
