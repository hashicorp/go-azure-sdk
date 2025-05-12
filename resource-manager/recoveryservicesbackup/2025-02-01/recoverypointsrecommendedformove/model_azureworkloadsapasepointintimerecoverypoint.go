package recoverypointsrecommendedformove

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadSAPAsePointInTimeRecoveryPoint{}

type AzureWorkloadSAPAsePointInTimeRecoveryPoint struct {
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties        *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	TimeRanges                     *[]PointInTimeRange                        `json:"timeRanges,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadSAPAsePointInTimeRecoveryPoint) RecoveryPoint() BaseRecoveryPointImpl {
	return BaseRecoveryPointImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadSAPAsePointInTimeRecoveryPoint{}

func (s AzureWorkloadSAPAsePointInTimeRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPAsePointInTimeRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPAsePointInTimeRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPAsePointInTimeRecoveryPoint: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSAPAsePointInTimeRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPAsePointInTimeRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
