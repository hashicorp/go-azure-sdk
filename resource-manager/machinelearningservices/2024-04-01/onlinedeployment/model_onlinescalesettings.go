package onlinedeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineScaleSettings interface {
	OnlineScaleSettings() BaseOnlineScaleSettingsImpl
}

var _ OnlineScaleSettings = BaseOnlineScaleSettingsImpl{}

type BaseOnlineScaleSettingsImpl struct {
	ScaleType ScaleType `json:"scaleType"`
}

func (s BaseOnlineScaleSettingsImpl) OnlineScaleSettings() BaseOnlineScaleSettingsImpl {
	return s
}

var _ OnlineScaleSettings = RawOnlineScaleSettingsImpl{}

// RawOnlineScaleSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawOnlineScaleSettingsImpl struct {
	onlineScaleSettings BaseOnlineScaleSettingsImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawOnlineScaleSettingsImpl) OnlineScaleSettings() BaseOnlineScaleSettingsImpl {
	return s.onlineScaleSettings
}

func (s RawOnlineScaleSettingsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalOnlineScaleSettingsImplementation(input []byte) (OnlineScaleSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OnlineScaleSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["scaleType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Default") {
		var out DefaultScaleSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DefaultScaleSettings: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TargetUtilization") {
		var out TargetUtilizationScaleSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetUtilizationScaleSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseOnlineScaleSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOnlineScaleSettingsImpl: %+v", err)
	}

	return RawOnlineScaleSettingsImpl{
		onlineScaleSettings: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
