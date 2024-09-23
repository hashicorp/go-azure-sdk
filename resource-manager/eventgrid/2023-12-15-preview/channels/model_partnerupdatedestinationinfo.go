package channels

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerUpdateDestinationInfo interface {
	PartnerUpdateDestinationInfo() BasePartnerUpdateDestinationInfoImpl
}

var _ PartnerUpdateDestinationInfo = BasePartnerUpdateDestinationInfoImpl{}

type BasePartnerUpdateDestinationInfoImpl struct {
	EndpointType PartnerEndpointType `json:"endpointType"`
}

func (s BasePartnerUpdateDestinationInfoImpl) PartnerUpdateDestinationInfo() BasePartnerUpdateDestinationInfoImpl {
	return s
}

var _ PartnerUpdateDestinationInfo = RawPartnerUpdateDestinationInfoImpl{}

// RawPartnerUpdateDestinationInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartnerUpdateDestinationInfoImpl struct {
	partnerUpdateDestinationInfo BasePartnerUpdateDestinationInfoImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawPartnerUpdateDestinationInfoImpl) PartnerUpdateDestinationInfo() BasePartnerUpdateDestinationInfoImpl {
	return s.partnerUpdateDestinationInfo
}

func UnmarshalPartnerUpdateDestinationInfoImplementation(input []byte) (PartnerUpdateDestinationInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerUpdateDestinationInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["endpointType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "WebHook") {
		var out WebhookUpdatePartnerDestinationInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebhookUpdatePartnerDestinationInfo: %+v", err)
		}
		return out, nil
	}

	var parent BasePartnerUpdateDestinationInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartnerUpdateDestinationInfoImpl: %+v", err)
	}

	return RawPartnerUpdateDestinationInfoImpl{
		partnerUpdateDestinationInfo: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
