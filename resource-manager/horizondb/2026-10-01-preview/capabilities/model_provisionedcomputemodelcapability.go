package capabilities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ComputeModelCapability = ProvisionedComputeModelCapability{}

type ProvisionedComputeModelCapability struct {
	RecommendedVCoreCount int64                   `json:"recommendedVCoreCount"`
	SupportedComputeSizes []ComputeSizeCapability `json:"supportedComputeSizes"`

	// Fields inherited from ComputeModelCapability

	Type PoolComputeModelType `json:"type"`
}

func (s ProvisionedComputeModelCapability) ComputeModelCapability() BaseComputeModelCapabilityImpl {
	return BaseComputeModelCapabilityImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ProvisionedComputeModelCapability{}

func (s ProvisionedComputeModelCapability) MarshalJSON() ([]byte, error) {
	type wrapper ProvisionedComputeModelCapability
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProvisionedComputeModelCapability: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProvisionedComputeModelCapability: %+v", err)
	}

	decoded["type"] = "Provisioned"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProvisionedComputeModelCapability: %+v", err)
	}

	return encoded, nil
}
