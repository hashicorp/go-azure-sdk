package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NCrossValidations interface {
	NCrossValidations() BaseNCrossValidationsImpl
}

var _ NCrossValidations = BaseNCrossValidationsImpl{}

type BaseNCrossValidationsImpl struct {
	Mode NCrossValidationsMode `json:"mode"`
}

func (s BaseNCrossValidationsImpl) NCrossValidations() BaseNCrossValidationsImpl {
	return s
}

var _ NCrossValidations = RawNCrossValidationsImpl{}

// RawNCrossValidationsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNCrossValidationsImpl struct {
	nCrossValidations BaseNCrossValidationsImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawNCrossValidationsImpl) NCrossValidations() BaseNCrossValidationsImpl {
	return s.nCrossValidations
}

func UnmarshalNCrossValidationsImplementation(input []byte) (NCrossValidations, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NCrossValidations into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Auto") {
		var out AutoNCrossValidations
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoNCrossValidations: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomNCrossValidations
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomNCrossValidations: %+v", err)
		}
		return out, nil
	}

	var parent BaseNCrossValidationsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNCrossValidationsImpl: %+v", err)
	}

	return RawNCrossValidationsImpl{
		nCrossValidations: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
