package fetchtieringcost

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FetchTieringCostInfoRequest = FetchTieringCostSavingsInfoForProtectedItemRequest{}

type FetchTieringCostSavingsInfoForProtectedItemRequest struct {
	ContainerName     string `json:"containerName"`
	ProtectedItemName string `json:"protectedItemName"`

	// Fields inherited from FetchTieringCostInfoRequest

	ObjectType     string                `json:"objectType"`
	SourceTierType RecoveryPointTierType `json:"sourceTierType"`
	TargetTierType RecoveryPointTierType `json:"targetTierType"`
}

func (s FetchTieringCostSavingsInfoForProtectedItemRequest) FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl {
	return BaseFetchTieringCostInfoRequestImpl{
		ObjectType:     s.ObjectType,
		SourceTierType: s.SourceTierType,
		TargetTierType: s.TargetTierType,
	}
}

var _ json.Marshaler = FetchTieringCostSavingsInfoForProtectedItemRequest{}

func (s FetchTieringCostSavingsInfoForProtectedItemRequest) MarshalJSON() ([]byte, error) {
	type wrapper FetchTieringCostSavingsInfoForProtectedItemRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FetchTieringCostSavingsInfoForProtectedItemRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FetchTieringCostSavingsInfoForProtectedItemRequest: %+v", err)
	}

	decoded["objectType"] = "FetchTieringCostSavingsInfoForProtectedItemRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FetchTieringCostSavingsInfoForProtectedItemRequest: %+v", err)
	}

	return encoded, nil
}
