package fetchtieringcost

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TieringCostInfo interface {
	TieringCostInfo() BaseTieringCostInfoImpl
}

var _ TieringCostInfo = BaseTieringCostInfoImpl{}

type BaseTieringCostInfoImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseTieringCostInfoImpl) TieringCostInfo() BaseTieringCostInfoImpl {
	return s
}

var _ TieringCostInfo = RawTieringCostInfoImpl{}

// RawTieringCostInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTieringCostInfoImpl struct {
	tieringCostInfo BaseTieringCostInfoImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawTieringCostInfoImpl) TieringCostInfo() BaseTieringCostInfoImpl {
	return s.tieringCostInfo
}

func UnmarshalTieringCostInfoImplementation(input []byte) (TieringCostInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TieringCostInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "TieringCostRehydrationInfo") {
		var out TieringCostRehydrationInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TieringCostRehydrationInfo: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TieringCostSavingInfo") {
		var out TieringCostSavingInfo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TieringCostSavingInfo: %+v", err)
		}
		return out, nil
	}

	var parent BaseTieringCostInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTieringCostInfoImpl: %+v", err)
	}

	return RawTieringCostInfoImpl{
		tieringCostInfo: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
