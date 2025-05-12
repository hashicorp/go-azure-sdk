package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = VirtualMachineTaskDetails{}

type VirtualMachineTaskDetails struct {
	SkippedReason       *string `json:"skippedReason,omitempty"`
	SkippedReasonString *string `json:"skippedReasonString,omitempty"`

	// Fields inherited from TaskTypeDetails

	InstanceType string `json:"instanceType"`
}

func (s VirtualMachineTaskDetails) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return BaseTaskTypeDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = VirtualMachineTaskDetails{}

func (s VirtualMachineTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper VirtualMachineTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualMachineTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualMachineTaskDetails: %+v", err)
	}

	decoded["instanceType"] = "VirtualMachineTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualMachineTaskDetails: %+v", err)
	}

	return encoded, nil
}
