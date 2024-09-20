package fetchtieringcost

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TieringCostInfo = TieringCostSavingInfo{}

type TieringCostSavingInfo struct {
	RetailSourceTierCostPerGBPerMonth float64 `json:"retailSourceTierCostPerGBPerMonth"`
	RetailTargetTierCostPerGBPerMonth float64 `json:"retailTargetTierCostPerGBPerMonth"`
	SourceTierSizeReductionInBytes    int64   `json:"sourceTierSizeReductionInBytes"`
	TargetTierSizeIncreaseInBytes     int64   `json:"targetTierSizeIncreaseInBytes"`

	// Fields inherited from TieringCostInfo

	ObjectType string `json:"objectType"`
}

func (s TieringCostSavingInfo) TieringCostInfo() BaseTieringCostInfoImpl {
	return BaseTieringCostInfoImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = TieringCostSavingInfo{}

func (s TieringCostSavingInfo) MarshalJSON() ([]byte, error) {
	type wrapper TieringCostSavingInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TieringCostSavingInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TieringCostSavingInfo: %+v", err)
	}

	decoded["objectType"] = "TieringCostSavingInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TieringCostSavingInfo: %+v", err)
	}

	return encoded, nil
}
