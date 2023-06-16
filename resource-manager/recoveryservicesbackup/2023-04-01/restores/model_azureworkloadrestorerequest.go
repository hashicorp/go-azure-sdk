package restores

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadRestoreRequest{}

type AzureWorkloadRestoreRequest struct {
	PropertyBag            *map[string]string `json:"propertyBag,omitempty"`
	RecoveryMode           *RecoveryMode      `json:"recoveryMode,omitempty"`
	RecoveryType           *RecoveryType      `json:"recoveryType,omitempty"`
	SourceResourceId       *string            `json:"sourceResourceId,omitempty"`
	TargetInfo             *TargetRestoreInfo `json:"targetInfo,omitempty"`
	TargetVirtualMachineId *string            `json:"targetVirtualMachineId,omitempty"`

	// Fields inherited from RestoreRequest
}

var _ json.Marshaler = AzureWorkloadRestoreRequest{}

func (s AzureWorkloadRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadRestoreRequest: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadRestoreRequest: %+v", err)
	}

	return encoded, nil
}
