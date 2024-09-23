package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Seasonality interface {
	Seasonality() BaseSeasonalityImpl
}

var _ Seasonality = BaseSeasonalityImpl{}

type BaseSeasonalityImpl struct {
	Mode SeasonalityMode `json:"mode"`
}

func (s BaseSeasonalityImpl) Seasonality() BaseSeasonalityImpl {
	return s
}

var _ Seasonality = RawSeasonalityImpl{}

// RawSeasonalityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSeasonalityImpl struct {
	seasonality BaseSeasonalityImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawSeasonalityImpl) Seasonality() BaseSeasonalityImpl {
	return s.seasonality
}

func UnmarshalSeasonalityImplementation(input []byte) (Seasonality, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Seasonality into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Auto") {
		var out AutoSeasonality
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoSeasonality: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomSeasonality
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomSeasonality: %+v", err)
		}
		return out, nil
	}

	var parent BaseSeasonalityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSeasonalityImpl: %+v", err)
	}

	return RawSeasonalityImpl{
		seasonality: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
