package dataprotections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureValidationResponseBase interface {
	FeatureValidationResponseBase() BaseFeatureValidationResponseBaseImpl
}

var _ FeatureValidationResponseBase = BaseFeatureValidationResponseBaseImpl{}

type BaseFeatureValidationResponseBaseImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseFeatureValidationResponseBaseImpl) FeatureValidationResponseBase() BaseFeatureValidationResponseBaseImpl {
	return s
}

var _ FeatureValidationResponseBase = RawFeatureValidationResponseBaseImpl{}

// RawFeatureValidationResponseBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFeatureValidationResponseBaseImpl struct {
	featureValidationResponseBase BaseFeatureValidationResponseBaseImpl
	Type                          string
	Values                        map[string]interface{}
}

func (s RawFeatureValidationResponseBaseImpl) FeatureValidationResponseBase() BaseFeatureValidationResponseBaseImpl {
	return s.featureValidationResponseBase
}

func UnmarshalFeatureValidationResponseBaseImplementation(input []byte) (FeatureValidationResponseBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureValidationResponseBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "FeatureValidationResponse") {
		var out FeatureValidationResponse
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureValidationResponse: %+v", err)
		}
		return out, nil
	}

	var parent BaseFeatureValidationResponseBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFeatureValidationResponseBaseImpl: %+v", err)
	}

	return RawFeatureValidationResponseBaseImpl{
		featureValidationResponseBase: parent,
		Type:                          value,
		Values:                        temp,
	}, nil

}
