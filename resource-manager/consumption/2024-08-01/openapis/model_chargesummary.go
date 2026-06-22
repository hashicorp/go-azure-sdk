package openapis

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargeSummary interface {
	ChargeSummary() BaseChargeSummaryImpl
}

var _ ChargeSummary = BaseChargeSummaryImpl{}

type BaseChargeSummaryImpl struct {
	ETag       *string                `json:"eTag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       ChargeSummaryKind      `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseChargeSummaryImpl) ChargeSummary() BaseChargeSummaryImpl {
	return s
}

var _ ChargeSummary = RawChargeSummaryImpl{}

// RawChargeSummaryImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawChargeSummaryImpl struct {
	chargeSummary BaseChargeSummaryImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawChargeSummaryImpl) ChargeSummary() BaseChargeSummaryImpl {
	return s.chargeSummary
}

func (s RawChargeSummaryImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalChargeSummaryImplementation(input []byte) (ChargeSummary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChargeSummary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
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
