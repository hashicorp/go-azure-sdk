package replicationprotectionintents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProtectionProfileCustomDetails = ExistingProtectionProfile{}

type ExistingProtectionProfile struct {
	ProtectionProfileId string `json:"protectionProfileId"`

	// Fields inherited from ProtectionProfileCustomDetails

	ResourceType string `json:"resourceType"`
}

func (s ExistingProtectionProfile) ProtectionProfileCustomDetails() BaseProtectionProfileCustomDetailsImpl {
	return BaseProtectionProfileCustomDetailsImpl{
		ResourceType: s.ResourceType,
	}
}

var _ json.Marshaler = ExistingProtectionProfile{}

func (s ExistingProtectionProfile) MarshalJSON() ([]byte, error) {
	type wrapper ExistingProtectionProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExistingProtectionProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExistingProtectionProfile: %+v", err)
	}

	decoded["resourceType"] = "Existing"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExistingProtectionProfile: %+v", err)
	}

	return encoded, nil
}
