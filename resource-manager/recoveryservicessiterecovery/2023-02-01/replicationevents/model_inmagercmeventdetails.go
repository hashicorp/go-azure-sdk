package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = InMageRcmEventDetails{}

type InMageRcmEventDetails struct {
	ApplianceName        *string `json:"applianceName,omitempty"`
	ComponentDisplayName *string `json:"componentDisplayName,omitempty"`
	FabricName           *string `json:"fabricName,omitempty"`
	JobId                *string `json:"jobId,omitempty"`
	LatestAgentVersion   *string `json:"latestAgentVersion,omitempty"`
	ProtectedItemName    *string `json:"protectedItemName,omitempty"`
	ServerType           *string `json:"serverType,omitempty"`
	VirtualMachineName   *string `json:"vmName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails
}

var _ json.Marshaler = InMageRcmEventDetails{}

func (s InMageRcmEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmEventDetails: %+v", err)
	}
	decoded["instanceType"] = "InMageRcm"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmEventDetails: %+v", err)
	}

	return encoded, nil
}
