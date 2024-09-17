package linkers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DryrunPrerequisiteResult interface {
	DryrunPrerequisiteResult() BaseDryrunPrerequisiteResultImpl
}

var _ DryrunPrerequisiteResult = BaseDryrunPrerequisiteResultImpl{}

type BaseDryrunPrerequisiteResultImpl struct {
	Type DryrunPrerequisiteResultType `json:"type"`
}

func (s BaseDryrunPrerequisiteResultImpl) DryrunPrerequisiteResult() BaseDryrunPrerequisiteResultImpl {
	return s
}

var _ DryrunPrerequisiteResult = RawDryrunPrerequisiteResultImpl{}

// RawDryrunPrerequisiteResultImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDryrunPrerequisiteResultImpl struct {
	dryrunPrerequisiteResult BaseDryrunPrerequisiteResultImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawDryrunPrerequisiteResultImpl) DryrunPrerequisiteResult() BaseDryrunPrerequisiteResultImpl {
	return s.dryrunPrerequisiteResult
}

func UnmarshalDryrunPrerequisiteResultImplementation(input []byte) (DryrunPrerequisiteResult, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DryrunPrerequisiteResult into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "basicError") {
		var out BasicErrorDryrunPrerequisiteResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BasicErrorDryrunPrerequisiteResult: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "permissionsMissing") {
		var out PermissionsMissingDryrunPrerequisiteResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PermissionsMissingDryrunPrerequisiteResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseDryrunPrerequisiteResultImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDryrunPrerequisiteResultImpl: %+v", err)
	}

	return RawDryrunPrerequisiteResultImpl{
		dryrunPrerequisiteResult: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
