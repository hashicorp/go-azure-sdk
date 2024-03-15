package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = DefenderForDevOpsGithubOffering{}

type DefenderForDevOpsGithubOffering struct {

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = DefenderForDevOpsGithubOffering{}

func (s DefenderForDevOpsGithubOffering) MarshalJSON() ([]byte, error) {
	type wrapper DefenderForDevOpsGithubOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DefenderForDevOpsGithubOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DefenderForDevOpsGithubOffering: %+v", err)
	}
	decoded["offeringType"] = "DefenderForDevOpsGithub"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DefenderForDevOpsGithubOffering: %+v", err)
	}

	return encoded, nil
}
