package fetchtieringcost

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FetchTieringCostInfoRequest = FetchTieringCostSavingsInfoForPolicyRequest{}

type FetchTieringCostSavingsInfoForPolicyRequest struct {
	PolicyName string `json:"policyName"`

	// Fields inherited from FetchTieringCostInfoRequest

	ObjectType     string                `json:"objectType"`
	SourceTierType RecoveryPointTierType `json:"sourceTierType"`
	TargetTierType RecoveryPointTierType `json:"targetTierType"`
}

func (s FetchTieringCostSavingsInfoForPolicyRequest) FetchTieringCostInfoRequest() BaseFetchTieringCostInfoRequestImpl {
	return BaseFetchTieringCostInfoRequestImpl{
		ObjectType:     s.ObjectType,
		SourceTierType: s.SourceTierType,
		TargetTierType: s.TargetTierType,
	}
}

var _ json.Marshaler = FetchTieringCostSavingsInfoForPolicyRequest{}

func (s FetchTieringCostSavingsInfoForPolicyRequest) MarshalJSON() ([]byte, error) {
	type wrapper FetchTieringCostSavingsInfoForPolicyRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FetchTieringCostSavingsInfoForPolicyRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FetchTieringCostSavingsInfoForPolicyRequest: %+v", err)
	}

	decoded["objectType"] = "FetchTieringCostSavingsInfoForPolicyRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FetchTieringCostSavingsInfoForPolicyRequest: %+v", err)
	}

	return encoded, nil
}
