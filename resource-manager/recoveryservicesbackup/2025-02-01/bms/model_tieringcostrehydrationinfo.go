package bms

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TieringCostInfo = TieringCostRehydrationInfo{}

type TieringCostRehydrationInfo struct {
	RehydrationSizeInBytes             int64   `json:"rehydrationSizeInBytes"`
	RetailRehydrationCostPerGBPerMonth float64 `json:"retailRehydrationCostPerGBPerMonth"`

	// Fields inherited from TieringCostInfo

	ObjectType string `json:"objectType"`
}

func (s TieringCostRehydrationInfo) TieringCostInfo() BaseTieringCostInfoImpl {
	return BaseTieringCostInfoImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = TieringCostRehydrationInfo{}

func (s TieringCostRehydrationInfo) MarshalJSON() ([]byte, error) {
	type wrapper TieringCostRehydrationInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TieringCostRehydrationInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TieringCostRehydrationInfo: %+v", err)
	}

	decoded["objectType"] = "TieringCostRehydrationInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TieringCostRehydrationInfo: %+v", err)
	}

	return encoded, nil
}
