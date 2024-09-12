package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomerInformationBase interface {
	BookingCustomerInformationBase() BaseBookingCustomerInformationBaseImpl
}

var _ BookingCustomerInformationBase = BaseBookingCustomerInformationBaseImpl{}

type BaseBookingCustomerInformationBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

func (s BaseBookingCustomerInformationBaseImpl) BookingCustomerInformationBase() BaseBookingCustomerInformationBaseImpl {
	return s
}

var _ BookingCustomerInformationBase = RawBookingCustomerInformationBaseImpl{}

// RawBookingCustomerInformationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBookingCustomerInformationBaseImpl struct {
	bookingCustomerInformationBase BaseBookingCustomerInformationBaseImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawBookingCustomerInformationBaseImpl) BookingCustomerInformationBase() BaseBookingCustomerInformationBaseImpl {
	return s.bookingCustomerInformationBase
}

func UnmarshalBookingCustomerInformationBaseImplementation(input []byte) (BookingCustomerInformationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BookingCustomerInformationBase into map[string]interface: %+v", err)
	}

	value, ok := temp["@odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.bookingCustomerInformation") {
		var out BookingCustomerInformation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BookingCustomerInformation: %+v", err)
		}
		return out, nil
	}

	var parent BaseBookingCustomerInformationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBookingCustomerInformationBaseImpl: %+v", err)
	}

	return RawBookingCustomerInformationBaseImpl{
		bookingCustomerInformationBase: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
