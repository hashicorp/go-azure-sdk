package channels

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PartnerDestinationInfo = WebhookPartnerDestinationInfo{}

type WebhookPartnerDestinationInfo struct {
	Properties *WebhookPartnerDestinationProperties `json:"properties,omitempty"`

	// Fields inherited from PartnerDestinationInfo

	AzureSubscriptionId       *string                      `json:"azureSubscriptionId,omitempty"`
	EndpointServiceContext    *string                      `json:"endpointServiceContext,omitempty"`
	EndpointType              PartnerEndpointType          `json:"endpointType"`
	Name                      *string                      `json:"name,omitempty"`
	ResourceGroupName         *string                      `json:"resourceGroupName,omitempty"`
	ResourceMoveChangeHistory *[]ResourceMoveChangeHistory `json:"resourceMoveChangeHistory,omitempty"`
}

func (s WebhookPartnerDestinationInfo) PartnerDestinationInfo() BasePartnerDestinationInfoImpl {
	return BasePartnerDestinationInfoImpl{
		AzureSubscriptionId:       s.AzureSubscriptionId,
		EndpointServiceContext:    s.EndpointServiceContext,
		EndpointType:              s.EndpointType,
		Name:                      s.Name,
		ResourceGroupName:         s.ResourceGroupName,
		ResourceMoveChangeHistory: s.ResourceMoveChangeHistory,
	}
}

var _ json.Marshaler = WebhookPartnerDestinationInfo{}

func (s WebhookPartnerDestinationInfo) MarshalJSON() ([]byte, error) {
	type wrapper WebhookPartnerDestinationInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebhookPartnerDestinationInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookPartnerDestinationInfo: %+v", err)
	}

	decoded["endpointType"] = "WebHook"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebhookPartnerDestinationInfo: %+v", err)
	}

	return encoded, nil
}
