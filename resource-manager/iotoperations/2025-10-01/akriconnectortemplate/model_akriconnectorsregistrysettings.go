package akriconnectortemplate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorsRegistrySettings interface {
	AkriConnectorsRegistrySettings() BaseAkriConnectorsRegistrySettingsImpl
}

var _ AkriConnectorsRegistrySettings = BaseAkriConnectorsRegistrySettingsImpl{}

type BaseAkriConnectorsRegistrySettingsImpl struct {
	RegistrySettingsType AkriConnectorsRegistrySettingsType `json:"registrySettingsType"`
}

func (s BaseAkriConnectorsRegistrySettingsImpl) AkriConnectorsRegistrySettings() BaseAkriConnectorsRegistrySettingsImpl {
	return s
}

var _ AkriConnectorsRegistrySettings = RawAkriConnectorsRegistrySettingsImpl{}

// RawAkriConnectorsRegistrySettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAkriConnectorsRegistrySettingsImpl struct {
	akriConnectorsRegistrySettings BaseAkriConnectorsRegistrySettingsImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawAkriConnectorsRegistrySettingsImpl) AkriConnectorsRegistrySettings() BaseAkriConnectorsRegistrySettingsImpl {
	return s.akriConnectorsRegistrySettings
}

func UnmarshalAkriConnectorsRegistrySettingsImplementation(input []byte) (AkriConnectorsRegistrySettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AkriConnectorsRegistrySettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["registrySettingsType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ContainerRegistry") {
		var out AkriConnectorsContainerRegistry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorsContainerRegistry: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RegistryEndpointRef") {
		var out AkriConnectorsRegistryEndpointRef
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AkriConnectorsRegistryEndpointRef: %+v", err)
		}
		return out, nil
	}

	var parent BaseAkriConnectorsRegistrySettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAkriConnectorsRegistrySettingsImpl: %+v", err)
	}

	return RawAkriConnectorsRegistrySettingsImpl{
		akriConnectorsRegistrySettings: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
