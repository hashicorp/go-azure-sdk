package connector

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

// RawDryrunParametersImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDryrunParametersImpl struct {
	dryrunParameters BaseDryrunParametersImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawDryrunParametersImpl) DryrunParameters() BaseDryrunParametersImpl {
	return s.dryrunParameters
}

func UnmarshalDryrunParametersImplementation(input []byte) (DryrunParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DryrunParameters into map[string]interface: %+v", err)
	}

	value, ok := temp["actionName"].(string)
	if !ok {
		return nil, nil
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
