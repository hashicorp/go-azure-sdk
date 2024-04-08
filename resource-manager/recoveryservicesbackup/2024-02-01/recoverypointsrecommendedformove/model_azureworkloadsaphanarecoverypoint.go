package recoverypointsrecommendedformove

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RecoveryPoint = AzureWorkloadSAPHanaRecoveryPoint{}

type AzureWorkloadSAPHanaRecoveryPoint struct {
	RecoveryPointMoveReadinessInfo *map[string]RecoveryPointMoveReadinessInfo `json:"recoveryPointMoveReadinessInfo,omitempty"`
	RecoveryPointProperties        *RecoveryPointProperties                   `json:"recoveryPointProperties,omitempty"`
	RecoveryPointTierDetails       *[]RecoveryPointTierInformationV2          `json:"recoveryPointTierDetails,omitempty"`
	RecoveryPointTimeInUTC         *string                                    `json:"recoveryPointTimeInUTC,omitempty"`
	Type                           *RestorePointType                          `json:"type,omitempty"`

	// Fields inherited from RecoveryPoint
}

var _ json.Marshaler = AzureWorkloadSAPHanaRecoveryPoint{}

func (s AzureWorkloadSAPHanaRecoveryPoint) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadSAPHanaRecoveryPoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadSAPHanaRecoveryPoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadSAPHanaRecoveryPoint: %+v", err)
	}
	decoded["objectType"] = "AzureWorkloadSAPHanaRecoveryPoint"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadSAPHanaRecoveryPoint: %+v", err)
	}

	return encoded, nil
}
