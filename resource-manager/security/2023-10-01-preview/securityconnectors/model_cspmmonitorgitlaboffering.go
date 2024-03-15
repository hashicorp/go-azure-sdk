package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = CspmMonitorGitLabOffering{}

type CspmMonitorGitLabOffering struct {

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = CspmMonitorGitLabOffering{}

func (s CspmMonitorGitLabOffering) MarshalJSON() ([]byte, error) {
	type wrapper CspmMonitorGitLabOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CspmMonitorGitLabOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CspmMonitorGitLabOffering: %+v", err)
	}
	decoded["offeringType"] = "CspmMonitorGitLab"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CspmMonitorGitLabOffering: %+v", err)
	}

	return encoded, nil
}
