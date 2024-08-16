package doordashdrive

import (
	"encoding/json"
	"fmt"
)

func (c *Client) CreateQuote(quoteRequest DeliveryRequest) (*DeliveryResponse, error) {
	resp, err := c.doRequest("POST", "/drive/v2/quotes", quoteRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quoteResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&quoteResponse)
	if err != nil {
		return nil, err
	}

	return &quoteResponse, nil
}

func (c *Client) AcceptQuote(externalDeliveryID string, acceptQuoteRequest AcceptQuoteRequest) (*DeliveryResponse, error) {
	path := fmt.Sprintf("/drive/v2/quotes/%s/accept", externalDeliveryID)
	resp, err := c.doRequest("POST", path, acceptQuoteRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var quoteResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&quoteResponse)
	if err != nil {
		return nil, err
	}

	return &quoteResponse, nil
}
