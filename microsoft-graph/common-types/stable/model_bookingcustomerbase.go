package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomerBase interface {
	Entity
	BookingCustomerBase() BaseBookingCustomerBaseImpl
}

var _ BookingCustomerBase = BaseBookingCustomerBaseImpl{}

type BaseBookingCustomerBaseImpl struct {

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseBookingCustomerBaseImpl) BookingCustomerBase() BaseBookingCustomerBaseImpl {
	return s
}

func (s BaseBookingCustomerBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ BookingCustomerBase = RawBookingCustomerBaseImpl{}

// RawBookingCustomerBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBookingCustomerBaseImpl struct {
	bookingCustomerBase BaseBookingCustomerBaseImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawBookingCustomerBaseImpl) BookingCustomerBase() BaseBookingCustomerBaseImpl {
	return s.bookingCustomerBase
}

func (s RawBookingCustomerBaseImpl) Entity() BaseEntityImpl {
	return s.bookingCustomerBase.Entity()
}

var _ json.Marshaler = BaseBookingCustomerBaseImpl{}

func (s BaseBookingCustomerBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseBookingCustomerBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseBookingCustomerBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseBookingCustomerBaseImpl: %+v", err)
	}

	decoded["@odata.type"] = "#microsoft.graph.bookingCustomerBase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseBookingCustomerBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalBookingCustomerBaseImplementation(input []byte) (BookingCustomerBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingCustomerBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomer") {
		var out BookingCustomer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomer: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingCustomerBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingCustomerBaseImpl: %+v", err)
	}

	return RawBookingCustomerBaseImpl{
		bookingCustomerBase: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
