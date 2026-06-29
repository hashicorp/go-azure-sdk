package sharesubscription

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceShareSynchronizationSetting interface {
	SourceShareSynchronizationSetting() BaseSourceShareSynchronizationSettingImpl
}

var _ SourceShareSynchronizationSetting = BaseSourceShareSynchronizationSettingImpl{}

type BaseSourceShareSynchronizationSettingImpl struct {
	Kind SourceShareSynchronizationSettingKind `json:"kind"`
}

func (s BaseSourceShareSynchronizationSettingImpl) SourceShareSynchronizationSetting() BaseSourceShareSynchronizationSettingImpl {
	return s
}

var _ SourceShareSynchronizationSetting = RawSourceShareSynchronizationSettingImpl{}

// RawSourceShareSynchronizationSettingImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawSourceShareSynchronizationSettingImpl struct {
	sourceShareSynchronizationSetting BaseSourceShareSynchronizationSettingImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawSourceShareSynchronizationSettingImpl) SourceShareSynchronizationSetting() BaseSourceShareSynchronizationSettingImpl {
	return s.sourceShareSynchronizationSetting
}

func (s RawSourceShareSynchronizationSettingImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalSourceShareSynchronizationSettingImplementation(input []byte) (SourceShareSynchronizationSetting, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SourceShareSynchronizationSetting into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ScheduleBased") {
		var out ScheduledSourceSynchronizationSetting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduledSourceSynchronizationSetting: %+v", err)
		}
		return out, nil
	}

	var parent BaseSourceShareSynchronizationSettingImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSourceShareSynchronizationSettingImpl: %+v", err)
	}

	return RawSourceShareSynchronizationSettingImpl{
		sourceShareSynchronizationSetting: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
