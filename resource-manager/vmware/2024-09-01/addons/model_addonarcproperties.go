package addons

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AddonProperties = AddonArcProperties{}

type AddonArcProperties struct {
	VCenter *string `json:"vCenter,omitempty"`

	// Fields inherited from AddonProperties

	AddonType         AddonType               `json:"addonType"`
	ProvisioningState *AddonProvisioningState `json:"provisioningState,omitempty"`
}

func (s AddonArcProperties) AddonProperties() BaseAddonPropertiesImpl {
	return BaseAddonPropertiesImpl{
		AddonType:         s.AddonType,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = AddonArcProperties{}

func (s AddonArcProperties) MarshalJSON() ([]byte, error) {
	type wrapper AddonArcProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddonArcProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonArcProperties: %+v", err)
	}

	decoded["addonType"] = "Arc"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddonArcProperties: %+v", err)
	}

	return encoded, nil
}
