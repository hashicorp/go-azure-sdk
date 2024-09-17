package channels

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerDestinationInfo interface {
	PartnerDestinationInfo() BasePartnerDestinationInfoImpl
}

var _ PartnerDestinationInfo = BasePartnerDestinationInfoImpl{}

type BasePartnerDestinationInfoImpl struct {
	AzureSubscriptionId       *string                      `json:"azureSubscriptionId,omitempty"`
	EndpointServiceContext    *string                      `json:"endpointServiceContext,omitempty"`
	EndpointType              PartnerEndpointType          `json:"endpointType"`
	Name                      *string                      `json:"name,omitempty"`
	ResourceGroupName         *string                      `json:"resourceGroupName,omitempty"`
	ResourceMoveChangeHistory *[]ResourceMoveChangeHistory `json:"resourceMoveChangeHistory,omitempty"`
}

func (s BasePartnerDestinationInfoImpl) PartnerDestinationInfo() BasePartnerDestinationInfoImpl {
	return s
}

var _ PartnerDestinationInfo = RawPartnerDestinationInfoImpl{}

// RawPartnerDestinationInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartnerDestinationInfoImpl struct {
	partnerDestinationInfo BasePartnerDestinationInfoImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawPartnerDestinationInfoImpl) PartnerDestinationInfo() BasePartnerDestinationInfoImpl {
	return s.partnerDestinationInfo
}

func UnmarshalPartnerDestinationInfoImplementation(input []byte) (PartnerDestinationInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartnerDestinationInfo into map[string]interface: %+v", err)
	}

	value, ok := temp["endpointType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "WebHook") {
		var out WebhookPartnerDestinationInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebhookPartnerDestinationInfo: %+v", err)
		}
		return out, nil
	}

	var parent BasePartnerDestinationInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartnerDestinationInfoImpl: %+v", err)
	}

	return RawPartnerDestinationInfoImpl{
		partnerDestinationInfo: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
