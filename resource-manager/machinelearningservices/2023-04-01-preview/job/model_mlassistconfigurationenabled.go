package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MLAssistConfiguration = MLAssistConfigurationEnabled{}

type MLAssistConfigurationEnabled struct {
	InferencingComputeBinding string `json:"inferencingComputeBinding"`
	TrainingComputeBinding    string `json:"trainingComputeBinding"`

	// Fields inherited from MLAssistConfiguration
}

var _ json.Marshaler = MLAssistConfigurationEnabled{}

func (s MLAssistConfigurationEnabled) MarshalJSON() ([]byte, error) {
	type wrapper MLAssistConfigurationEnabled
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MLAssistConfigurationEnabled: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MLAssistConfigurationEnabled: %+v", err)
	}
	decoded["mlAssist"] = "Enabled"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MLAssistConfigurationEnabled: %+v", err)
	}

	return encoded, nil
}
