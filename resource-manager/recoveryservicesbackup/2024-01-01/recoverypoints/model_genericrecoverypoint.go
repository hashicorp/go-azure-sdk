package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = GenericRecoveryPoint{}

type GenericRecoveryPoint struct {
	FriendlyName                *string                  `json:"friendlyName,omitempty"`
	RecoveryPointAdditionalInfo *string                  `json:"recoveryPointAdditionalInfo,omitempty"`
	RecoveryPointProperties     *RecoveryPointProperties `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTime           *string                  `json:"recoveryPointTime,omitempty"`
	RecoveryPointType           *string                  `json:"recoveryPointType,omitempty"`

	// Fields inherited from RecoveryPoint
}

var _ json.Marshaler = GenericRecoveryPoint{}

func (s GenericRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper GenericRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GenericRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GenericRecoveryPoint: %+v", err)
	}
	decoded["objectType"] = "GenericRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GenericRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
