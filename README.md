# DoorDash Drive

Golang client for the [DoorDash Drive](https://developer.doordash.com/en-US/api/drive/) (v2) API.

It is still in active development and testing: it has not yet been used in production. Version will be updated to 1.0.0+ once it has been used successfully in production.

This SDK will not cover the Drive (classic) API nor the Marketplace API.

## Installation

`go get github.com/rareplanet1/doordashdrive-sdk-go`

## Usage

All [DoorDash Drive](https://developer.doordash.com/en-US/api/drive/) (v2) API endpoints have been implemented to create & accept quotes, create deliveries, and manage businesses and stores.

```go
client := doordashdrive.NewClient("your_developer_id", "your_key_id", "your_signing_secret")

deliveryRequest := doordashdrive.DeliveryRequest{
	ExternalDeliveryID:     externalDeliveryId,
	OrderFulfillmentMethod: "standard",
	ActionIfUndeliverable:  "return_to_pickup",
	PickupAddress:          "901 Market Street 6th Floor San Francisco, CA 94103",
	PickupBusinessName:     "Wells Fargo SF Downtown",
	PickupInstructions:     "Collect from desk in lobby",
	PickupPhoneNumber:      "+16505555555",
	PickupWindow: &doordashdrive.TimeWindow{
		StartTime: StrToTime("2024-08-15T13:00:00Z"),
		EndTime:   StrToTime("2024-08-15T13:30:00Z"),
	},
	DropoffAddress:      "901 Market Street 6th Floor San Francisco, CA 94103",
	DropoffBusinessName: "The Avery Condominium",
	DropoffInstructions: "Deliver to reception",
	DropoffPhoneNumber:  "+61400123456",
	DropoffWindow: &doordashdrive.TimeWindow{
		StartTime: StrToTime("2024-08-15T14:00:00Z"),
		EndTime:   StrToTime("2024-08-15T14:30:00Z"),
	},
	OrderValue: 10000,
}

deliveryResponse, err := client.CreateDelivery(deliveryRequest)
if err != nil {
	if apiErr, ok := err.(*doordashdrive.APIError); ok {
		for _, fieldError := range apiErr.FieldErrors {
			fmt.Printf("Field error: %s: %s\n", fieldError.Field, fieldError.Error)
		}
		log.Fatalf("API error: %v", apiErr)
	}
	log.Fatalf("Error creating delivery: %v", err)
}

fmt.Printf("Delivery created: %+v\n", deliveryResponse)
```

See examples folder for more.


## Contributions

Pull requests welcome.


