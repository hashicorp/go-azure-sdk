package securityconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CloudOffering = CspmMonitorAzureDevOpsOffering{}

type CspmMonitorAzureDevOpsOffering struct {

	// Fields inherited from CloudOffering
	Description *string `json:"description,omitempty"`
}

var _ json.Marshaler = CspmMonitorAzureDevOpsOffering{}

func (s CspmMonitorAzureDevOpsOffering) MarshalJSON() ([]byte, error) {
	type wrapper CspmMonitorAzureDevOpsOffering
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CspmMonitorAzureDevOpsOffering: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CspmMonitorAzureDevOpsOffering: %+v", err)
	}
	decoded["offeringType"] = "CspmMonitorAzureDevOps"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CspmMonitorAzureDevOpsOffering: %+v", err)
	}

	return encoded, nil
}
