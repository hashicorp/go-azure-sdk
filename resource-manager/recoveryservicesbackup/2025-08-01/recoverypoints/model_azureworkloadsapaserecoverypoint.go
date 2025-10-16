package recoverypoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadSAPAseRecoveryPoint{}

type AzureWorkloadSAPAseRecoveryPoint struct {
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties        *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint

	ObjectType string `json:"objectType"`
}

func (s AzureWorkloadSAPAseRecoveryPoint) RecoveryPoint() BaseRecoveryPointImpl {
	return BaseRecoveryPointImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = AzureWorkloadSAPAseRecoveryPoint{}

func (s AzureWorkloadSAPAseRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPAseRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPAseRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPAseRecoveryPoint: %+v", err)
	}

	decoded["objectType"] = "AzureWorkloadSAPAseRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPAseRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
