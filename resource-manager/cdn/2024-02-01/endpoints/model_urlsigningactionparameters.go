package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleActionParameters = UrlSigningActionParameters{}

type UrlSigningActionParameters struct {
	Algorithm             *Algorithm                   `json:"algorithm,omitempty"`
	ParameterNameOverride *[]UrlSigningParamIdentifier `json:"parameterNameOverride,omitempty"`

	// Fields inherited from DeliveryRuleActionParameters
}

var _ json.Marshaler = UrlSigningActionParameters{}

func (s UrlSigningActionParameters) MarshalJSON() ([]byte, error) {
	type wrapper UrlSigningActionParameters
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlSigningActionParameters: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlSigningActionParameters: %+v", err)
	}
	decoded["typeName"] = "DeliveryRuleUrlSigningActionParameters"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlSigningActionParameters: %+v", err)
	}

	return encoded, nil
}
