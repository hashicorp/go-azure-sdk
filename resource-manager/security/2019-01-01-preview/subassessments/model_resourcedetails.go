package subassessments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceDetails interface {
	ResourceDetails() BaseResourceDetailsImpl
}

var _ ResourceDetails = BaseResourceDetailsImpl{}

type BaseResourceDetailsImpl struct {
	Source Source `json:"source"`
}

func (s BaseResourceDetailsImpl) ResourceDetails() BaseResourceDetailsImpl {
	return s
}

var _ ResourceDetails = RawResourceDetailsImpl{}

// RawResourceDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawResourceDetailsImpl struct {
	resourceDetails BaseResourceDetailsImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawResourceDetailsImpl) ResourceDetails() BaseResourceDetailsImpl {
	return s.resourceDetails
}

func UnmarshalResourceDetailsImplementation(input []byte) (ResourceDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["source"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Azure") {
		var out AzureResourceDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureResourceDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "OnPremise") {
		var out OnPremiseResourceDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnPremiseResourceDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseResourceDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseResourceDetailsImpl: %+v", err)
	}

	return RawResourceDetailsImpl{
		resourceDetails: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
