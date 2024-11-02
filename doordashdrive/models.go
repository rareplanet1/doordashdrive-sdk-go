package doordashdrive

import "time"

type DeliveryRequest struct {
	ExternalDeliveryID                 string                `json:"external_delivery_id"`
	Locale                             string                `json:"locale,omitempty"`
	OrderFulfillmentMethod             string                `json:"order_fulfillment_method"`
	OriginFacilityID                   string                `json:"origin_facility_id,omitempty"`
	PickupAddress                      string                `json:"pickup_address,omitempty"`
	PickupBusinessName                 string                `json:"pickup_business_name,omitempty"`
	PickupPhoneNumber                  string                `json:"pickup_phone_number,omitempty"`
	PickupInstructions                 string                `json:"pickup_instructions,omitempty"`
	PickupReferenceTag                 string                `json:"pickup_reference_tag,omitempty"`
	PickupExternalBusinessID           string                `json:"pickup_external_business_id,omitempty"`
	PickupExternalStoreID              string                `json:"pickup_external_store_id,omitempty"`
	PickupVerificationMetadata         *VerificationMetadata `json:"pickup_verification_metadata,omitempty"`
	DropoffAddress                     string                `json:"dropoff_address,omitempty"`
	DropoffBusinessName                string                `json:"dropoff_business_name,omitempty"`
	DropoffLocation                    *Location             `json:"dropoff_location,omitempty"`
	DropoffPhoneNumber                 string                `json:"dropoff_phone_number,omitempty"`
	DropoffInstructions                string                `json:"dropoff_instructions,omitempty"`
	DropoffContactGivenName            string                `json:"dropoff_contact_given_name,omitempty"`
	DropoffContactFamilyName           string                `json:"dropoff_contact_family_name,omitempty"`
	DropoffContactSendNotifications    bool                  `json:"dropoff_contact_send_notifications"`
	DropoffOptions                     *DropoffOptions       `json:"dropoff_options,omitempty"`
	DropoffAddressComponents           *AddressComponents    `json:"dropoff_address_components,omitempty"`
	DropoffPinCodeVerificationMetadata *PinCodeVerification  `json:"dropoff_pin_code_verification_metadata,omitempty"`
	ShoppingOptions                    *ShoppingOptions      `json:"shopping_options,omitempty"`
	OrderValue                         int                   `json:"order_value,omitempty"`
	Items                              []Item                `json:"items,omitempty"`
	PickupTime                         *time.Time            `json:"pickup_time,omitempty"`
	DropoffTime                        *time.Time            `json:"dropoff_time,omitempty"`
	PickupWindow                       *TimeWindow           `json:"pickup_window,omitempty"`
	DropoffWindow                      *TimeWindow           `json:"dropoff_window,omitempty"`
	ContactlessDropoff                 bool                  `json:"contactless_dropoff,omitempty"`
	ActionIfUndeliverable              string                `json:"action_if_undeliverable"`
	Tip                                int                   `json:"tip,omitempty"`
	OrderContains                      *OrderContains        `json:"order_contains,omitempty"`
	DasherAllowedVehicles              []string              `json:"dasher_allowed_vehicles,omitempty"`
	DropoffRequiresSignature           bool                  `json:"dropoff_requires_signature,omitempty"`
	PromotionID                        string                `json:"promotion_id,omitempty"`
	DropoffCashOnDelivery              int                   `json:"dropoff_cash_on_delivery,omitempty"`
	OrderRouteType                     string                `json:"order_route_type,omitempty"`
	OrderRouteItems                    []string              `json:"order_route_items,omitempty"`
}

type VerificationMetadata struct {
	VerificationType   string `json:"verification_type,omitempty"`
	VerificationCode   string `json:"verification_code,omitempty"`
	VerificationFormat string `json:"verification_format,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type DropoffOptions struct {
	Signature       string `json:"signature,omitempty"`
	IDVerification  string `json:"id_verification,omitempty"`
	ProofOfDelivery string `json:"proof_of_delivery,omitempty"`
	CateringSetup   string `json:"catering_setup,omitempty"`
}

type AddressComponents struct {
	StreetAddress string `json:"street_address,omitempty"` // Used to denote the street number and street name; multiple lines should be concatenated with a comma
	SubPremise    string `json:"sub_premise,omitempty"`    // Used to denote the unit number, flat number, suite number, or apartment number
	City          string `json:"city,omitempty"`           // US & CA only
	SubLocality   string `json:"sub_locality,omitempty"`   // NZ only
	Locality      string `json:"locality,omitempty"`       // AU & NZ only
	State         string `json:"state,omitempty"`          // US only
	Province      string `json:"province,omitempty"`       // CA only
	Territory     string `json:"territory,omitempty"`      // AU only
	ZipCode       string `json:"zip_code,omitempty"`       // US only
	PostalCode    string `json:"postal_code,omitempty"`    // CA, AU, NZ only
	Country       string `json:"country,omitempty"`
}

type PinCodeVerification struct {
	PinCodeType  string `json:"pin_code_type,omitempty"`
	PinCodeValue string `json:"pin_code_value,omitempty"`
}

type ShoppingOptions struct {
	PaymentMethod               string     `json:"payment_method,omitempty"`
	PaymentBarcode              string     `json:"payment_barcode,omitempty"`
	PaymentGiftCards            []string   `json:"payment_gift_cards,omitempty"`
	ReadyForPickupBy            *time.Time `json:"ready_for_pickup_by,omitempty"`
	DropoffContactLoyaltyNumber string     `json:"dropoff_contact_loyalty_number,omitempty"`
}

type Item struct {
	Name               string             `json:"name"`
	Description        string             `json:"description,omitempty"`
	Quantity           int                `json:"quantity"`
	ExternalID         string             `json:"external_id"`
	ExternalInstanceID int                `json:"external_instance_id,omitempty"`
	Volume             float64            `json:"volume,omitempty"`
	Weight             float64            `json:"weight,omitempty"`
	Length             float64            `json:"length,omitempty"`
	Width              float64            `json:"width,omitempty"`
	Height             float64            `json:"height,omitempty"`
	Price              int                `json:"price"`
	Barcode            int64              `json:"barcode,omitempty"`
	ItemOptions        *ItemOptions       `json:"item_options,omitempty"`
	AdjustmentDetails  *AdjustmentDetails `json:"adjustment_details,omitempty"`
}

type ItemOptions struct {
	SubstituteItemIDs      []string `json:"substitute_item_ids,omitempty"`
	WeightUnit             string   `json:"weight_unit,omitempty"`
	SubstitutionPreference string   `json:"substitution_preference,omitempty"`
}

type AdjustmentDetails struct {
	AdditionSource string `json:"addition_source,omitempty"`
}

type TimeWindow struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type OrderContains struct {
	Alcohol                    bool `json:"alcohol,omitempty"`
	PharmacyItems              bool `json:"pharmacy_items,omitempty"`
	AgeRestrictedPharmacyItems bool `json:"age_restricted_pharmacy_items,omitempty"`
	Tobacco                    bool `json:"tobacco,omitempty"`
	Hemp                       bool `json:"hemp,omitempty"`
	OTC                        bool `json:"otc,omitempty"`
}

type DeliveryResponse struct {
	ExternalDeliveryID                 string               `json:"external_delivery_id"`
	Locale                             string               `json:"locale,omitempty"`
	OrderFulfillmentMethod             string               `json:"order_fulfillment_method,omitempty"`
	OriginFacilityID                   string               `json:"origin_facility_id,omitempty"`
	PickupAddress                      string               `json:"pickup_address,omitempty"`
	PickupBusinessName                 string               `json:"pickup_business_name,omitempty"`
	PickupPhoneNumber                  string               `json:"pickup_phone_number,omitempty"`
	PickupInstructions                 string               `json:"pickup_instructions,omitempty"`
	PickupReferenceTag                 string               `json:"pickup_reference_tag,omitempty"`
	PickupExternalBusinessID           string               `json:"pickup_external_business_id,omitempty"`
	PickupExternalStoreID              string               `json:"pickup_external_store_id,omitempty"`
	PickupVerificationMetadata         VerificationMetadata `json:"pickup_verification_metadata,omitempty"`
	DropoffAddress                     string               `json:"dropoff_address,omitempty"`
	DropoffBusinessName                string               `json:"dropoff_business_name,omitempty"`
	DropoffLocation                    Location             `json:"dropoff_location,omitempty"`
	DropoffPhoneNumber                 string               `json:"dropoff_phone_number,omitempty"`
	DropoffInstructions                string               `json:"dropoff_instructions,omitempty"`
	DropoffContactGivenName            string               `json:"dropoff_contact_given_name,omitempty"`
	DropoffContactFamilyName           string               `json:"dropoff_contact_family_name,omitempty"`
	DropoffContactSendNotifications    bool                 `json:"dropoff_contact_send_notifications,omitempty"`
	DropoffOptions                     DropoffOptions       `json:"dropoff_options,omitempty"`
	DropoffAddressComponents           AddressComponents    `json:"dropoff_address_components,omitempty"`
	DropoffPinCodeVerificationMetadata PinCodeVerification  `json:"dropoff_pin_code_verification_metadata,omitempty"`
	OrderValue                         int                  `json:"order_value,omitempty"`
	Currency                           string               `json:"currency,omitempty"`
	Items                              []Item               `json:"items,omitempty"`
	ShoppingOptions                    ShoppingOptions      `json:"shopping_options,omitempty"`
	DeliveryStatus                     string               `json:"delivery_status,omitempty"`
	CancellationReason                 string               `json:"cancellation_reason,omitempty"`
	UpdatedAt                          *time.Time           `json:"updated_at,omitempty"`
	PickupTimeEstimated                *time.Time           `json:"pickup_time_estimated,omitempty"`
	PickupTimeActual                   *time.Time           `json:"pickup_time_actual,omitempty"`
	DropoffTimeEstimated               *time.Time           `json:"dropoff_time_estimated,omitempty"`
	DropoffTimeActual                  *time.Time           `json:"dropoff_time_actual,omitempty"`
	ReturnTimeEstimated                *time.Time           `json:"return_time_estimated,omitempty"`
	ReturnTimeActual                   *time.Time           `json:"return_time_actual,omitempty"`
	ReturnAddress                      string               `json:"return_address,omitempty"`
	Fee                                int                  `json:"fee,omitempty"`
	FeeComponents                      []FeeComponent       `json:"fee_components,omitempty"`
	Tax                                int                  `json:"tax,omitempty"`
	TaxComponents                      []TaxComponent       `json:"tax_components,omitempty"`
	SupportReference                   string               `json:"support_reference,omitempty"`
	TrackingURL                        string               `json:"tracking_url,omitempty"`
	DropoffVerificationImageURL        string               `json:"dropoff_verification_image_url,omitempty"`
	PickupVerificationImageURL         string               `json:"pickup_verification_image_url,omitempty"`
	DropoffSignatureImageURL           string               `json:"dropoff_signature_image_url,omitempty"`
	ShippingLabel                      ShippingLabel        `json:"shipping_label,omitempty"`
	DroppedItems                       []DroppedItem        `json:"dropped_items,omitempty"`
	ContactlessDropoff                 bool                 `json:"contactless_dropoff,omitempty"`
	ActionIfUndeliverable              string               `json:"action_if_undeliverable,omitempty"`
	Tip                                int                  `json:"tip,omitempty"`
	OrderContains                      OrderContains        `json:"order_contains,omitempty"`
	DasherAllowedVehicles              []string             `json:"dasher_allowed_vehicles,omitempty"`
	DropoffRequiresSignature           bool                 `json:"dropoff_requires_signature,omitempty"`
	PromotionID                        string               `json:"promotion_id,omitempty"`
	DropoffCashOnDelivery              int                  `json:"dropoff_cash_on_delivery,omitempty"`
	OrderRouteType                     string               `json:"order_route_type,omitempty"`
	OrderRouteItems                    []string             `json:"order_route_items,omitempty"`
	DasherID                           int                  `json:"dasher_id,omitempty"`
	DasherName                         string               `json:"dasher_name,omitempty"`
	DasherDropoffPhoneNumber           string               `json:"dasher_dropoff_phone_number,omitempty"`
	DasherPickupPhoneNumber            string               `json:"dasher_pickup_phone_number,omitempty"`
	DasherLocation                     Location             `json:"dasher_location,omitempty"`
	DasherVehicleMake                  string               `json:"dasher_vehicle_make,omitempty"`
	DasherVehicleModel                 string               `json:"dasher_vehicle_model,omitempty"`
	DasherVehicleYear                  string               `json:"dasher_vehicle_year,omitempty"`
}

type FeeComponent struct {
	Type   string `json:"type,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type TaxComponent struct {
	Type   string `json:"type,omitempty"`
	Amount int    `json:"amount,omitempty"`
}

type ShippingLabel struct {
	LabelFormat  string `json:"label_format,omitempty"`
	LabelSize    string `json:"label_size,omitempty"`
	PrintDensity string `json:"print_density,omitempty"`
	LabelString  string `json:"label_string,omitempty"`
}

type DroppedItem struct {
	ExternalID string `json:"external_id,omitempty"`
	Type       string `json:"type,omitempty"`
	Reason     string `json:"reason,omitempty"`
}

type AcceptQuoteRequest struct {
	Tip                int    `json:"tip,omitempty"`
	DropoffPhoneNumber string `json:"dropoff_phone_number,omitempty"`
}
