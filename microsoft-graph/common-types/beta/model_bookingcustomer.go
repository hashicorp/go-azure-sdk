package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BookingPerson = BookingCustomer{}

type BookingCustomer struct {
	// Addresses associated with the customer, including home, business and other addresses.
	Addresses *[]PhysicalAddress `json:"addresses,omitempty"`

	// The date, time, and timezone when the customer was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The date, time, and timezone when the customer was last updated.
	LastUpdatedDateTime nullable.Type[string] `json:"lastUpdatedDateTime,omitempty"`

	// Phone numbers associated with the customer, including home, business and mobile numbers.
	Phones *[]Phone `json:"phones,omitempty"`

	// Fields inherited from BookingPerson

	// The email address of the person.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Fields inherited from BookingNamedEntity

	// A name for the derived entity, which interfaces with customers.
	DisplayName *string `json:"displayName,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BookingCustomer) BookingPerson() BaseBookingPersonImpl {
	return BaseBookingPersonImpl{
		EmailAddress: s.EmailAddress,
		DisplayName:  s.DisplayName,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s BookingCustomer) BookingNamedEntity() BaseBookingNamedEntityImpl {
	return BaseBookingNamedEntityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s BookingCustomer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BookingCustomer{}

func (s BookingCustomer) MarshalJSON() ([]byte, error) {
	type wrapper BookingCustomer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BookingCustomer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingCustomer: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.bookingCustomer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BookingCustomer: %+v", err)
	}

	return encoded, nil
}
