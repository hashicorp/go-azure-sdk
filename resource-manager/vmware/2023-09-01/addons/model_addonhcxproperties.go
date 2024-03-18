package addons

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AddonProperties = AddonHcxProperties{}

type AddonHcxProperties struct {
	Offer string `json:"offer"`

	// Fields inherited from AddonProperties
	ProvisioningState *AddonProvisioningState `json:"provisioningState,omitempty"`
}

var _ json.Marshaler = AddonHcxProperties{}

func (s AddonHcxProperties) MarshalJSON() ([]byte, error) {
	type wrapper AddonHcxProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AddonHcxProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonHcxProperties: %+v", err)
	}
	decoded["addonType"] = "HCX"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AddonHcxProperties: %+v", err)
	}

	return encoded, nil
}
