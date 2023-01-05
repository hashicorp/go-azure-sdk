package dppfeaturesupport

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureValidationRequestBase interface {
}

func unmarshalFeatureValidationRequestBaseImplementation(input []byte) (FeatureValidationRequestBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureValidationRequestBase into map[string]interface: %+v", err)
	}

	value, ok := temp["objectType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "FeatureValidationRequest") {
		var out FeatureValidationRequest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureValidationRequest: %+v", err)
		}
		return out, nil
	}

	type RawFeatureValidationRequestBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawFeatureValidationRequestBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
