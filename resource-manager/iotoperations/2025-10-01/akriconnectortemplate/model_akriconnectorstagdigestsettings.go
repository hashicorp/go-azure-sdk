package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsTagDigestSettings interface {
	AkriConnectorsTagDigestSettings() BaseAkriConnectorsTagDigestSettingsImpl
}

var _ AkriConnectorsTagDigestSettings = BaseAkriConnectorsTagDigestSettingsImpl{}

type BaseAkriConnectorsTagDigestSettingsImpl struct {
	TagDigestType AkriConnectorsTagDigestType `json:"tagDigestType"`
}

func (s BaseAkriConnectorsTagDigestSettingsImpl) AkriConnectorsTagDigestSettings() BaseAkriConnectorsTagDigestSettingsImpl {
	return s
}

var _ AkriConnectorsTagDigestSettings = RawAkriConnectorsTagDigestSettingsImpl{}

// RawAkriConnectorsTagDigestSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAkriConnectorsTagDigestSettingsImpl struct {
	akriConnectorsTagDigestSettings BaseAkriConnectorsTagDigestSettingsImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawAkriConnectorsTagDigestSettingsImpl) AkriConnectorsTagDigestSettings() BaseAkriConnectorsTagDigestSettingsImpl {
	return s.akriConnectorsTagDigestSettings
}

func (s RawAkriConnectorsTagDigestSettingsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAkriConnectorsTagDigestSettingsImplementation(input []byte) (AkriConnectorsTagDigestSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsTagDigestSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["tagDigestType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Digest") {
		var out AkriConnectorsDigest
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorsDigest: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Tag") {
		var out AkriConnectorsTag
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorsTag: %+v", err)
		}
		return out, nil
	}

	var parent BaseAkriConnectorsTagDigestSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAkriConnectorsTagDigestSettingsImpl: %+v", err)
	}

	return RawAkriConnectorsTagDigestSettingsImpl{
		akriConnectorsTagDigestSettings: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
