package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedDiskReplicationProviderSpecificSettings interface {
	SharedDiskReplicationProviderSpecificSettings() BaseSharedDiskReplicationProviderSpecificSettingsImpl
}

var _ SharedDiskReplicationProviderSpecificSettings = BaseSharedDiskReplicationProviderSpecificSettingsImpl{}

type BaseSharedDiskReplicationProviderSpecificSettingsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseSharedDiskReplicationProviderSpecificSettingsImpl) SharedDiskReplicationProviderSpecificSettings() BaseSharedDiskReplicationProviderSpecificSettingsImpl {
	return s
}

var _ SharedDiskReplicationProviderSpecificSettings = RawSharedDiskReplicationProviderSpecificSettingsImpl{}

// RawSharedDiskReplicationProviderSpecificSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSharedDiskReplicationProviderSpecificSettingsImpl struct {
	sharedDiskReplicationProviderSpecificSettings BaseSharedDiskReplicationProviderSpecificSettingsImpl
	Type                                          string
	Values                                        map[string]interface{}
}

func (s RawSharedDiskReplicationProviderSpecificSettingsImpl) SharedDiskReplicationProviderSpecificSettings() BaseSharedDiskReplicationProviderSpecificSettingsImpl {
	return s.sharedDiskReplicationProviderSpecificSettings
}

func UnmarshalSharedDiskReplicationProviderSpecificSettingsImplementation(input []byte) (SharedDiskReplicationProviderSpecificSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SharedDiskReplicationProviderSpecificSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out A2ASharedDiskReplicationDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2ASharedDiskReplicationDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseSharedDiskReplicationProviderSpecificSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSharedDiskReplicationProviderSpecificSettingsImpl: %+v", err)
	}

	return RawSharedDiskReplicationProviderSpecificSettingsImpl{
		sharedDiskReplicationProviderSpecificSettings: parent,
		Type:   value,
		Values: temp,
	}, nil

}
