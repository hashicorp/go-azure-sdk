package addons

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AddonProperties = AddonVrProperties{}

type AddonVrProperties struct {
	VrsCount int64 `json:"vrsCount"`

	// Fields inherited from AddonProperties

	AddonType         AddonType               `json:"addonType"`
	ProvisioningState *AddonProvisioningState `json:"provisioningState,omitempty"`
}

func (s AddonVrProperties) AddonProperties() BaseAddonPropertiesImpl {
	return BaseAddonPropertiesImpl{
		AddonType:         s.AddonType,
		ProvisioningState: s.ProvisioningState,
	}
}

var _ json.Marshaler = AddonVrProperties{}

func (s AddonVrProperties) MarshalJSON() ([]byte, error) {
	type wrapper AddonVrProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddonVrProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonVrProperties: %+v", err)
	}

	decoded["addonType"] = "VR"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddonVrProperties: %+v", err)
	}

	return encoded, nil
}
