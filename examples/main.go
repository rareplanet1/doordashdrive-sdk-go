package main

import (
	"fmt"
	"log"

	"github.com/rareplanet1/doordashdrive-sdk-go/doordashdrive"
)

func main() {
	client := doordashdrive.NewClient("your_developer_id", "your_key_id", "your_signing_secret")

	externalDeliveryId := "Delivery-123"
	runQuoteRequest(client, externalDeliveryId)
	// runAcceptQuoteRequest(client, externalDeliveryId)
	// runDeliveryRequest(client, externalDeliveryId)
	// runGetDelivery(client, externalDeliveryId)

	// externalBusinessId := "Business-123"
	// externalStoreId := "Store-123"
	// runCreateBusiness(client, externalBusinessId)
	// runGetBusinesses(client)
	// runCreateStore(client, externalBusinessId, externalStoreId)
	// runGetStore(client, externalBusinessId, externalStoreId)
	// runGetBusinesses(client)
	// runGetStores(client, externalBusinessId)
}

func runQuoteRequest(client *doordashdrive.Client, externalDeliveryId string) {
	quoteRequest := doordashdrive.DeliveryRequest{
		ExternalDeliveryID:     externalDeliveryId,
		OrderFulfillmentMethod: "standard",
		ActionIfUndeliverable:  "return_to_pickup",
		PickupAddress:          "901 Market Street 6th Floor San Francisco, CA 94103",
		PickupBusinessName:     "Wells Fargo SF Downtown",
		PickupInstructions:     "Collect from desk in lobby",
		PickupPhoneNumber:      "+16505555555",
		// PickupTime:             StrToTime("2024-08-15T13:00:00Z"),
		PickupWindow: &doordashdrive.TimeWindow{
			StartTime: StrToTime("2024-08-15T13:00:00Z"),
			EndTime:   StrToTime("2024-08-15T13:30:00Z"),
		},
		DropoffAddress:      "901 Market Street 6th Floor San Francisco, CA 94103",
		DropoffBusinessName: "The Avery Condominium",
		DropoffInstructions: "Deliver to reception",
		DropoffPhoneNumber:  "+61400123456",
		// DropoffTime:         StrToTime("2024-08-15T14:00:00Z"),
		DropoffWindow: &doordashdrive.TimeWindow{
			StartTime: StrToTime("2024-08-15T14:00:00Z"),
			EndTime:   StrToTime("2024-08-15T14:30:00Z"),
		},
		OrderValue: 10000,
	}

	quoteResponse, err := client.CreateQuote(quoteRequest)
	if err != nil {
		if apiErr, ok := err.(*doordashdrive.APIError); ok {
			for _, fieldError := range apiErr.FieldErrors {
				fmt.Printf("Field error: %s: %s\n", fieldError.Field, fieldError.Error)
			}
			log.Fatalf("API error: %v", apiErr)
		}
		log.Fatalf("Error creating quote: %v", err)
	}

	fmt.Printf("Quote created: %+v\n", quoteResponse)
}

func runAcceptQuoteRequest(client *doordashdrive.Client, externalDeliveryId string) {
	acceptQuoteRequest := doordashdrive.AcceptQuoteRequest{
		Tip:                500,
		DropoffPhoneNumber: "+61400123456",
	}

	quoteResponse, err := client.AcceptQuote(externalDeliveryId, acceptQuoteRequest)
	if err != nil {
		if apiErr, ok := err.(*doordashdrive.APIError); ok {
			for _, fieldError := range apiErr.FieldErrors {
				fmt.Printf("Field error: %s: %s\n", fieldError.Field, fieldError.Error)
			}
			log.Fatalf("API error: %v", apiErr)
		}
		log.Fatalf("Error accepting quote: %v", err)
	}

	fmt.Printf("Quote accepted: %+v\n", quoteResponse)
}

func runDeliveryRequest(client *doordashdrive.Client, externalDeliveryId string) {
	deliveryRequest := doordashdrive.DeliveryRequest{
		ExternalDeliveryID:     externalDeliveryId,
		OrderFulfillmentMethod: "standard",
		ActionIfUndeliverable:  "return_to_pickup",
		PickupAddress:          "901 Market Street 6th Floor San Francisco, CA 94103",
		PickupBusinessName:     "Wells Fargo SF Downtown",
		PickupInstructions:     "Collect from desk in lobby",
		PickupPhoneNumber:      "+16505555555",
		// PickupTime:             StrToTime("2024-08-15T13:00:00Z"),
		PickupWindow: &doordashdrive.TimeWindow{
			StartTime: StrToTime("2024-08-15T13:00:00Z"),
			EndTime:   StrToTime("2024-08-15T13:30:00Z"),
		},
		DropoffAddress:      "901 Market Street 6th Floor San Francisco, CA 94103",
		DropoffBusinessName: "The Avery Condominium",
		DropoffInstructions: "Deliver to reception",
		DropoffPhoneNumber:  "+61400123456",
		// DropoffTime:         StrToTime("2024-08-15T14:00:00Z"),
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
}

func runGetDelivery(client *doordashdrive.Client, externalDeliveryId string) {
	delivery, err := client.GetDelivery(externalDeliveryId)
	if err != nil {
		log.Fatalf("Error getting delivery: %v", err)
	}

	fmt.Printf("Delivery: %+v\n", delivery)
}

func runCancelDelivery(client *doordashdrive.Client, externalDeliveryId string) {
	delivery, err := client.CancelDelivery(externalDeliveryId)
	if err != nil {
		log.Fatalf("Error canceling delivery: %v", err)
	}

	fmt.Printf("Delivery canceled: %+v\n", delivery)
}
func runCreateBusiness(client *doordashdrive.Client, externalBusinessId string) {
	businessRequest := doordashdrive.BusinessRequest{
		ExternalBusinessID: externalBusinessId,
		Name:               "My Business",
		Description:        "My Business Description",
		ActivationStatus:   "active",
	}

	businessResponse, err := client.CreateBusiness(businessRequest)
	if err != nil {
		if apiErr, ok := err.(*doordashdrive.APIError); ok {
			for _, fieldError := range apiErr.FieldErrors {
				fmt.Printf("Field error: %s: %s\n", fieldError.Field, fieldError.Error)
			}
			log.Fatalf("API error: %v", apiErr)
		}
		log.Fatalf("Error creating business: %v", err)
	}

	fmt.Printf("Business created: %+v\n", businessResponse)
}

func runCreateStore(client *doordashdrive.Client, externalBusinessId, externalStoreId string) {
	storeRequest := doordashdrive.StoreRequest{
		ExternalBusinessID: externalBusinessId,
		ExternalStoreID:    externalStoreId,
		Name:               "My Store",
		PhoneNumber:        "+16505555555",
		Address:            "123 Main St, San Francisco, CA 94103",
	}

	storeResponse, err := client.CreateStore(externalBusinessId, storeRequest)
	if err != nil {
		if apiErr, ok := err.(*doordashdrive.APIError); ok {
			for _, fieldError := range apiErr.FieldErrors {
				fmt.Printf("Field error: %s: %s\n", fieldError.Field, fieldError.Error)
			}
			log.Fatalf("API error: %v", apiErr)
		}
		log.Fatalf("Error creating store: %v", err)
	}

	fmt.Printf("Store created: %+v\n", storeResponse)
}

func runGetBusinesses(client *doordashdrive.Client) {
	businesses, err := client.GetBusinesses()
	if err != nil {
		log.Fatalf("Error getting businesses: %v", err)
	}

	fmt.Printf("Businesses: %+v\n", businesses)
}

func runGetBusiness(client *doordashdrive.Client, externalBusinessId string) {
	business, err := client.GetBusiness(externalBusinessId)
	if err != nil {
		log.Fatalf("Error getting business: %v", err)
	}

	fmt.Printf("Business: %+v\n", business)
}

func runGetStores(client *doordashdrive.Client, externalBusinessId string) {
	stores, err := client.GetStores(externalBusinessId)
	if err != nil {
		log.Fatalf("Error getting stores: %v", err)
	}

	fmt.Printf("Stores: %+v\n", stores)
}

func runGetStore(client *doordashdrive.Client, externalBusinessId, externalStoreId string) {
	store, err := client.GetStore(externalBusinessId, externalStoreId)
	if err != nil {
		log.Fatalf("Error getting store: %v", err)
	}

	fmt.Printf("Store: %+v\n", store)
}
