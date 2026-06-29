package linkers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DryrunParameters interface {
	DryrunParameters() BaseDryrunParametersImpl
}

var _ DryrunParameters = BaseDryrunParametersImpl{}

type BaseDryrunParametersImpl struct {
	ActionName DryrunActionName `json:"actionName"`
}

func (s BaseDryrunParametersImpl) DryrunParameters() BaseDryrunParametersImpl {
	return s
}

var _ DryrunParameters = RawDryrunParametersImpl{}

// RawDryrunParametersImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDryrunParametersImpl struct {
	dryrunParameters BaseDryrunParametersImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawDryrunParametersImpl) DryrunParameters() BaseDryrunParametersImpl {
	return s.dryrunParameters
}

func (s RawDryrunParametersImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDryrunParametersImplementation(input []byte) (DryrunParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DryrunParameters into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["actionName"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "createOrUpdate") {
		var out CreateOrUpdateDryrunParameters
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CreateOrUpdateDryrunParameters: %+v", err)
		}
		return out, nil
	}

	var parent BaseDryrunParametersImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDryrunParametersImpl: %+v", err)
	}

	return RawDryrunParametersImpl{
		dryrunParameters: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
