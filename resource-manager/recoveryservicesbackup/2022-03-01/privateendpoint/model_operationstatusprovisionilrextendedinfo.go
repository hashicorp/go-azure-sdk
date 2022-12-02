package privateendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OperationStatusExtendedInfo = OperationStatusProvisionILRExtendedInfo{}

type OperationStatusProvisionILRExtendedInfo struct {
	RecoveryTarget *InstantItemRecoveryTarget `json:"recoveryTarget,omitempty"`

	// Fields inherited from OperationStatusExtendedInfo
}

var _ json.Marshaler = OperationStatusProvisionILRExtendedInfo{}

func (s OperationStatusProvisionILRExtendedInfo) MarshalJSON() ([]byte, error) {
	type wrapper OperationStatusProvisionILRExtendedInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationStatusProvisionILRExtendedInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusProvisionILRExtendedInfo: %+v", err)
	}
	decoded["objectType"] = "OperationStatusProvisionILRExtendedInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationStatusProvisionILRExtendedInfo: %+v", err)
	}

	return encoded, nil
}
