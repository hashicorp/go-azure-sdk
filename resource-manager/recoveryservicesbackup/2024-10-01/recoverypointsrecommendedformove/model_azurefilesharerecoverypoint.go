package recoverypointsrecommendedformove

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureFileShareRecoveryPoint{}

type AzureFileShareRecoveryPoint struct {
	FileShareSnapshotUri     *string                         `json:"fileShareSnapshotUri,omitempty"`
	RecoveryPointProperties  *RecoveryPointProperties        `json:"recoveryPointProperties,omitempty"`
	RecoveryPointSizeInGB    *int64                          `json:"recoveryPointSizeInGB,omitempty"`
	RecoveryPointTierDetails *[]RecoveryPointTierInformation `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTime        *string                         `json:"recoveryPointTime,omitempty"`
	RecoveryPointType        *string                         `json:"recoveryPointType,omitempty"`

	// Fields inherited from RecoveryPoint

	ObjectType string `json:"objectType"`
}

func (s AzureFileShareRecoveryPoint) RecoveryPoint() BaseRecoveryPointImpl {
	return BaseRecoveryPointImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureFileShareRecoveryPoint{}

func (s AzureFileShareRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureFileShareRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureFileShareRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureFileShareRecoveryPoint: %+v", err)
	}

	decoded["objectType"] = "AzureFileShareRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureFileShareRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
