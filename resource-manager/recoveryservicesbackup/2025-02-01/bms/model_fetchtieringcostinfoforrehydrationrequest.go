package bms

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FetchTieringCostInfoRequest = FetchTieringCostInfoForRehydrationRequest{}

type FetchTieringCostInfoForRehydrationRequest struct {
	ContainerName       string              `json:"containerName"`
	ProtectedItemName   string              `json:"protectedItemName"`
	RecoveryPointId     string              `json:"recoveryPointId"`
	RehydrationPriority RehydrationPriority `json:"rehydrationPriority"`

	// Fields inherited from FetchTieringCostInfoRequest

	ObjectType     string                `json:"objectType"`
	SourceTierType RecoveryPointTierType `json:"sourceTierType"`
	TargetTierType RecoveryPointTierType `json:"targetTierType"`
}

func (s FetchTieringCostInfoForRehydrationRequest) FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl {
	return BaseFetchTieringCostInfoRequestImpl{
		ObjectType:     s.ObjectType,
		SourceTierType: s.SourceTierType,
		TargetTierType: s.TargetTierType,
	}
}

var _ json.Marshaler = FetchTieringCostInfoForRehydrationRequest{}

func (s FetchTieringCostInfoForRehydrationRequest) MarshalJSON() ([]byte, error) {
	type wrapper FetchTieringCostInfoForRehydrationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FetchTieringCostInfoForRehydrationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FetchTieringCostInfoForRehydrationRequest: %+v", err)
	}

	decoded["objectType"] = "FetchTieringCostInfoForRehydrationRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FetchTieringCostInfoForRehydrationRequest: %+v", err)
	}

	return encoded, nil
}
