package fetchtieringcost

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FetchTieringCostInfoRequest = FetchTieringCostSavingsInfoForVaultRequest{}

type FetchTieringCostSavingsInfoForVaultRequest struct {

	// Fields inherited from FetchTieringCostInfoRequest
	SourceTierType RecoveryPointTierType `json:"sourceTierType"`
	TargetTierType RecoveryPointTierType `json:"targetTierType"`
}

var _ json.Marshaler = FetchTieringCostSavingsInfoForVaultRequest{}

func (s FetchTieringCostSavingsInfoForVaultRequest) MarshalJSON() ([]byte, error) {
	type wrapper FetchTieringCostSavingsInfoForVaultRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FetchTieringCostSavingsInfoForVaultRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FetchTieringCostSavingsInfoForVaultRequest: %+v", err)
	}
	decoded["objectType"] = "FetchTieringCostSavingsInfoForVaultRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FetchTieringCostSavingsInfoForVaultRequest: %+v", err)
	}

	return encoded, nil
}
