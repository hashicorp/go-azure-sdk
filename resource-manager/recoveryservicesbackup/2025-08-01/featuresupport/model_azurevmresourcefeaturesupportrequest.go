package featuresupport

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FeatureSupportRequest = AzureVMResourceFeatureSupportRequest{}

type AzureVMResourceFeatureSupportRequest struct {
	VMSize *string `json:"vmSize,omitempty"`
	VMSku  *string `json:"vmSku,omitempty"`

	// Fields inherited from FeatureSupportRequest

	FeatureType string `json:"featureType"`
}

func (s AzureVMResourceFeatureSupportRequest) FeatureSupportRequest() BaseFeatureSupportRequestImpl {
	return BaseFeatureSupportRequestImpl{
		FeatureType: s.FeatureType,
	}
}

var _ json.Marshaler = AzureVMResourceFeatureSupportRequest{}

func (s AzureVMResourceFeatureSupportRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureVMResourceFeatureSupportRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureVMResourceFeatureSupportRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureVMResourceFeatureSupportRequest: %+v", err)
	}

	decoded["featureType"] = "AzureVMResourceBackup"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureVMResourceFeatureSupportRequest: %+v", err)
	}

	return encoded, nil
}
