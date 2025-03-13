package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitorComputeIdentityBase = AmlTokenComputeIdentity{}

type AmlTokenComputeIdentity struct {

	// Fields inherited from MonitorComputeIdentityBase

	ComputeIdentityType MonitorComputeIdentityType `json:"computeIdentityType"`
}

func (s AmlTokenComputeIdentity) MonitorComputeIdentityBase() BaseMonitorComputeIdentityBaseImpl {
	return BaseMonitorComputeIdentityBaseImpl{
		ComputeIdentityType: s.ComputeIdentityType,
	}
}

var _ json.Marshaler = AmlTokenComputeIdentity{}

func (s AmlTokenComputeIdentity) MarshalJSON() ([]byte, error) {
	type wrapper AmlTokenComputeIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AmlTokenComputeIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AmlTokenComputeIdentity: %+v", err)
	}

	decoded["computeIdentityType"] = "AmlToken"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AmlTokenComputeIdentity: %+v", err)
	}

	return encoded, nil
}
