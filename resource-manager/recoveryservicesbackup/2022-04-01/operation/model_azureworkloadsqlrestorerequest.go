package operation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RestoreRequest = AzureWorkloadSQLRestoreRequest{}

type AzureWorkloadSQLRestoreRequest struct {
	AlternateDirectoryPaths          *[]SQLDataDirectoryMapping `json:"alternateDirectoryPaths,omitempty"`
	IsNonRecoverable                 *bool                      `json:"isNonRecoverable,omitempty"`
	PropertyBag                      *map[string]string         `json:"propertyBag,omitempty"`
	RecoveryMode                     *RecoveryMode              `json:"recoveryMode,omitempty"`
	RecoveryType                     *RecoveryType              `json:"recoveryType,omitempty"`
	ShouldUseAlternateTargetLocation *bool                      `json:"shouldUseAlternateTargetLocation,omitempty"`
	SourceResourceId                 *string                    `json:"sourceResourceId,omitempty"`
	TargetInfo                       *TargetRestoreInfo         `json:"targetInfo,omitempty"`
	TargetVirtualMachineId           *string                    `json:"targetVirtualMachineId,omitempty"`

	// Fields inherited from RestoreRequest
}

var _ json.Marshaler = AzureWorkloadSQLRestoreRequest{}

func (s AzureWorkloadSQLRestoreRequest) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLRestoreRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLRestoreRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLRestoreRequest: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadSQLRestoreRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLRestoreRequest: %+v", err)
	}

	return encoded, nil
}
