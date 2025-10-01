package bms

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FeatureSupportRequest = AzureBackupGoalFeatureSupportRequest{}

type AzureBackupGoalFeatureSupportRequest struct {

	// Fields inherited from FeatureSupportRequest

	FeatureType string `json:"featureType"`
}

func (s AzureBackupGoalFeatureSupportRequest) FeatureSupportRequest() BaseFeatureSupportRequestImpl {
	return BaseFeatureSupportRequestImpl{
		FeatureType: s.FeatureType,
	}
}

var _ json.Marshaler = AzureBackupGoalFeatureSupportRequest{}

func (s AzureBackupGoalFeatureSupportRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureBackupGoalFeatureSupportRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureBackupGoalFeatureSupportRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureBackupGoalFeatureSupportRequest: %+v", err)
	}

	decoded["featureType"] = "AzureBackupGoals"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureBackupGoalFeatureSupportRequest: %+v", err)
	}

	return encoded, nil
}
