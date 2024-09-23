package resources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Facet interface {
	Facet() BaseFacetImpl
}

var _ Facet = BaseFacetImpl{}

type BaseFacetImpl struct {
	Expression string `json:"expression"`
	ResultType string `json:"resultType"`
}

func (s BaseFacetImpl) Facet() BaseFacetImpl {
	return s
}

var _ Facet = RawFacetImpl{}

// RawFacetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFacetImpl struct {
	facet  BaseFacetImpl
	Type   string
	Values map[string]interface{}
}

func (s RawFacetImpl) Facet() BaseFacetImpl {
	return s.facet
}

func UnmarshalFacetImplementation(input []byte) (Facet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Facet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["resultType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "FacetError") {
		var out FacetError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FacetError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FacetResult") {
		var out FacetResult
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FacetResult: %+v", err)
		}
		return out, nil
	}

	var parent BaseFacetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFacetImpl: %+v", err)
	}

	return RawFacetImpl{
		facet:  parent,
		Type:   value,
		Values: temp,
	}, nil

}
