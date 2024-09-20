package charges

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargeSummary interface {
	ChargeSummary() BaseChargeSummaryImpl
}

var _ ChargeSummary = BaseChargeSummaryImpl{}

type BaseChargeSummaryImpl struct {
	ETag *string            `json:"eTag,omitempty"`
	Id   *string            `json:"id,omitempty"`
	Kind ChargeSummaryKind  `json:"kind"`
	Name *string            `json:"name,omitempty"`
	Tags *map[string]string `json:"tags,omitempty"`
	Type *string            `json:"type,omitempty"`
}

func (s BaseChargeSummaryImpl) ChargeSummary() BaseChargeSummaryImpl {
	return s
}

var _ ChargeSummary = RawChargeSummaryImpl{}

// RawChargeSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawChargeSummaryImpl struct {
	chargeSummary BaseChargeSummaryImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawChargeSummaryImpl) ChargeSummary() BaseChargeSummaryImpl {
	return s.chargeSummary
}

func UnmarshalChargeSummaryImplementation(input []byte) (ChargeSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChargeSummary into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "legacy") {
		var out LegacyChargeSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacyChargeSummary: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "modern") {
		var out ModernChargeSummary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernChargeSummary: %+v", err)
		}
		return out, nil
	}

	var parent BaseChargeSummaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChargeSummaryImpl: %+v", err)
	}

	return RawChargeSummaryImpl{
		chargeSummary: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
