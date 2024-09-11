package rules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleActionParameters = UrlRedirectActionParameters{}

type UrlRedirectActionParameters struct {
	CustomFragment      *string              `json:"customFragment,omitempty"`
	CustomHostname      *string              `json:"customHostname,omitempty"`
	CustomPath          *string              `json:"customPath,omitempty"`
	CustomQueryString   *string              `json:"customQueryString,omitempty"`
	DestinationProtocol *DestinationProtocol `json:"destinationProtocol,omitempty"`
	RedirectType        RedirectType         `json:"redirectType"`

	// Fields inherited from DeliveryRuleActionParameters
}

var _ json.Marshaler = UrlRedirectActionParameters{}

func (s UrlRedirectActionParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlRedirectActionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlRedirectActionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlRedirectActionParameters: %+v", err)
	}
	decoded["typeName"] = "DeliveryRuleUrlRedirectActionParameters"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlRedirectActionParameters: %+v", err)
	}

	return encoded, nil
}
