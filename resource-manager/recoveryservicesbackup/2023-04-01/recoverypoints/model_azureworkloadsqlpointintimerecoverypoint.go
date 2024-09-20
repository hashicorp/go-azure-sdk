package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadSQLPointInTimeRecoveryPoint{}

type AzureWorkloadSQLPointInTimeRecoveryPoint struct {
	ExtendedInfo                   *AzureWorkloadSQLRecoveryPointExtendedInfo `json:"extendedInfo,omitempty"`
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties        *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	TimeRanges                     *[]PointInTimeRange                        `json:"timeRanges,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadSQLPointInTimeRecoveryPoint) RecoveryPoint() BaseRecoveryPointImpl {
	return BaseRecoveryPointImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadSQLPointInTimeRecoveryPoint{}

func (s AzureWorkloadSQLPointInTimeRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSQLPointInTimeRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSQLPointInTimeRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSQLPointInTimeRecoveryPoint: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSQLPointInTimeRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSQLPointInTimeRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
