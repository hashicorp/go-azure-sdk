package addons

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AddonProperties = AddonSrmProperties{}

type AddonSrmProperties struct {
	LicenseKey *string `json:"licenseKey,omitempty"`

	// Fields inherited from AddonProperties
	ProvisioningState *AddonProvisioningState `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = AddonSrmProperties{}

func (s AddonSrmProperties) MarshalJSON() ([]byte, error) {
	type wrapper AddonSrmProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddonSrmProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonSrmProperties: %+v", err)
	}
	decoded["addonType"] = "SRM"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddonSrmProperties: %+v", err)
	}

	return encoded, nil
}
