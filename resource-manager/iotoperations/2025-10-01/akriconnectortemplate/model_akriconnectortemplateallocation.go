package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateAllocation interface {
	AkriConnectorTemplateAllocation() BaseAkriConnectorTemplateAllocationImpl
}

var _ AkriConnectorTemplateAllocation = BaseAkriConnectorTemplateAllocationImpl{}

type BaseAkriConnectorTemplateAllocationImpl struct {
	Policy AkriConnectorTemplateAllocationPolicy `json:"policy"`
}

func (s BaseAkriConnectorTemplateAllocationImpl) AkriConnectorTemplateAllocation() BaseAkriConnectorTemplateAllocationImpl {
	return s
}

var _ AkriConnectorTemplateAllocation = RawAkriConnectorTemplateAllocationImpl{}

// RawAkriConnectorTemplateAllocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAkriConnectorTemplateAllocationImpl struct {
	akriConnectorTemplateAllocation BaseAkriConnectorTemplateAllocationImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawAkriConnectorTemplateAllocationImpl) AkriConnectorTemplateAllocation() BaseAkriConnectorTemplateAllocationImpl {
	return s.akriConnectorTemplateAllocation
}

func UnmarshalAkriConnectorTemplateAllocationImplementation(input []byte) (AkriConnectorTemplateAllocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorTemplateAllocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["policy"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Bucketized") {
		var out AkriConnectorTemplateBucketizedAllocation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorTemplateBucketizedAllocation: %+v", err)
		}
		return out, nil
	}

	var parent BaseAkriConnectorTemplateAllocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAkriConnectorTemplateAllocationImpl: %+v", err)
	}

	return RawAkriConnectorTemplateAllocationImpl{
		akriConnectorTemplateAllocation: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
