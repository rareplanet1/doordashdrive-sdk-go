package doordashdrive

import (
	"encoding/json"
	"fmt"
)

// CreateDelivery creates a new delivery
func (c *Client) CreateDelivery(deliveryRequest DeliveryRequest) (*DeliveryResponse, error) {
	resp, err := c.doRequest("POST", "/drive/v2/deliveries", deliveryRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deliveryResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&deliveryResponse)
	if err != nil {
		return nil, err
	}

	return &deliveryResponse, nil
}

// GetDelivery retrieves a delivery by its external ID
func (c *Client) GetDelivery(externalDeliveryID string) (*DeliveryResponse, error) {
	path := fmt.Sprintf("/drive/v2/deliveries/%s", externalDeliveryID)
	resp, err := c.doRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deliveryResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&deliveryResponse)
	if err != nil {
		return nil, err
	}

	return &deliveryResponse, nil
}

// UpdateDelivery updates an existing delivery
func (c *Client) UpdateDelivery(externalDeliveryID string, updateRequest DeliveryRequest) (*DeliveryResponse, error) {
	path := fmt.Sprintf("/drive/v2/deliveries/%s", externalDeliveryID)
	resp, err := c.doRequest("PATCH", path, updateRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deliveryResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&deliveryResponse)
	if err != nil {
		return nil, err
	}

	return &deliveryResponse, nil
}

// CancelDelivery cancels an existing delivery
func (c *Client) CancelDelivery(externalDeliveryID string) (*DeliveryResponse, error) {
	path := fmt.Sprintf("/drive/v2/deliveries/%s/cancel", externalDeliveryID)
	resp, err := c.doRequest("PUT", path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var deliveryResponse DeliveryResponse
	err = json.NewDecoder(resp.Body).Decode(&deliveryResponse)
	if err != nil {
		return nil, err
	}

	return &deliveryResponse, nil

}
