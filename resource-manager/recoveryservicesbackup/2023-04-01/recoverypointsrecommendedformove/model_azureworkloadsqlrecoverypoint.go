package recoverypointsrecommendedformove

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadSQLRecoveryPoint{}

type AzureWorkloadSQLRecoveryPoint struct {
	ExtendedInfo                   *AzureWorkloadSQLRecoveryPointExtendedInfo `json:"extendedInfo,omitempty"`
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties        *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint
}

var _ json.Marshaler = AzureWorkloadSQLRecoveryPoint{}

func (s AzureWorkloadSQLRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLRecoveryPoint: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadSQLRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
