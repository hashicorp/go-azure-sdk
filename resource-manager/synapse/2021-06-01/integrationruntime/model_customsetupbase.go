package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSetupBase interface {
	CustomSetupBase() BaseCustomSetupBaseImpl
}

var _ CustomSetupBase = BaseCustomSetupBaseImpl{}

type BaseCustomSetupBaseImpl struct {
	Type string `json:"type"`
}

func (s BaseCustomSetupBaseImpl) CustomSetupBase() BaseCustomSetupBaseImpl {
	return s
}

var _ CustomSetupBase = RawCustomSetupBaseImpl{}

// RawCustomSetupBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomSetupBaseImpl struct {
	customSetupBase BaseCustomSetupBaseImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawCustomSetupBaseImpl) CustomSetupBase() BaseCustomSetupBaseImpl {
	return s.customSetupBase
}

func UnmarshalCustomSetupBaseImplementation(input []byte) (CustomSetupBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomSetupBase into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	var parent BaseCustomSetupBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomSetupBaseImpl: %+v", err)
	}

	return RawCustomSetupBaseImpl{
		customSetupBase: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
