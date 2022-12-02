package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadPointInTimeRecoveryPoint{}

type AzureWorkloadPointInTimeRecoveryPoint struct {
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	TimeRanges                     *[]PointInTimeRange                        `json:"timeRanges,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint
}

var _ json.Marshaler = AzureWorkloadPointInTimeRecoveryPoint{}

func (s AzureWorkloadPointInTimeRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadPointInTimeRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadPointInTimeRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadPointInTimeRecoveryPoint: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadPointInTimeRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadPointInTimeRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
