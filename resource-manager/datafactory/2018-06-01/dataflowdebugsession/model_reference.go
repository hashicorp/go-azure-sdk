package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reference interface {
	Reference() BaseReferenceImpl
}

var _ Reference = BaseReferenceImpl{}

type BaseReferenceImpl struct {
	Type string `json:"type"`
}

func (s BaseReferenceImpl) Reference() BaseReferenceImpl {
	return s
}

var _ Reference = RawReferenceImpl{}

// RawReferenceImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawReferenceImpl struct {
	reference BaseReferenceImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawReferenceImpl) Reference() BaseReferenceImpl {
	return s.reference
}

func (s RawReferenceImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalReferenceImplementation(input []byte) (Reference, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Reference into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "IntegrationRuntimeReference") {
		var out IntegrationRuntimeReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IntegrationRuntimeReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "LinkedServiceReference") {
		var out LinkedServiceReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LinkedServiceReference: %+v", err)
		}
		return out, nil
	}

	var parent BaseReferenceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseReferenceImpl: %+v", err)
	}

	return RawReferenceImpl{
		reference: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
