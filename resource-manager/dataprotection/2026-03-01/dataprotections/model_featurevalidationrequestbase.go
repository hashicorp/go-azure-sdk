package dataprotections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureValidationRequestBase interface {
	FeatureValidationRequestBase() BaseFeatureValidationRequestBaseImpl
}

var _ FeatureValidationRequestBase = BaseFeatureValidationRequestBaseImpl{}

type BaseFeatureValidationRequestBaseImpl struct {
	ObjectType string `json:"objectType"`
}

func (s BaseFeatureValidationRequestBaseImpl) FeatureValidationRequestBase() BaseFeatureValidationRequestBaseImpl {
	return s
}

var _ FeatureValidationRequestBase = RawFeatureValidationRequestBaseImpl{}

// RawFeatureValidationRequestBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawFeatureValidationRequestBaseImpl struct {
	featureValidationRequestBase BaseFeatureValidationRequestBaseImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawFeatureValidationRequestBaseImpl) FeatureValidationRequestBase() BaseFeatureValidationRequestBaseImpl {
	return s.featureValidationRequestBase
}

func (s RawFeatureValidationRequestBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalFeatureValidationRequestBaseImplementation(input []byte) (FeatureValidationRequestBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureValidationRequestBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["objectType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "FeatureValidationRequest") {
		var out FeatureValidationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureValidationRequest: %+v", err)
		}
		return out, nil
	}

	var parent BaseFeatureValidationRequestBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFeatureValidationRequestBaseImpl: %+v", err)
	}

	return RawFeatureValidationRequestBaseImpl{
		featureValidationRequestBase: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
