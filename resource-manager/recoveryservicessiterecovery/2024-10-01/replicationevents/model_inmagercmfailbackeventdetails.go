package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = InMageRcmFailbackEventDetails{}

type InMageRcmFailbackEventDetails struct {
	ApplianceName        *string `json:"applianceName,omitempty"`
	ComponentDisplayName *string `json:"componentDisplayName,omitempty"`
	ProtectedItemName    *string `json:"protectedItemName,omitempty"`
	ServerType           *string `json:"serverType,omitempty"`
	VirtualMachineName   *string `json:"vmName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails

	InstanceType string `json:"instanceType"`
}

func (s InMageRcmFailbackEventDetails) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return BaseEventProviderSpecificDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = InMageRcmFailbackEventDetails{}

func (s InMageRcmFailbackEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper InMageRcmFailbackEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InMageRcmFailbackEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InMageRcmFailbackEventDetails: %+v", err)
	}

	decoded["instanceType"] = "InMageRcmFailback"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InMageRcmFailbackEventDetails: %+v", err)
	}

	return encoded, nil
}
