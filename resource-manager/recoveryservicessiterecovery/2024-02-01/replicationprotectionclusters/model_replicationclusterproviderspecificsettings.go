package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationClusterProviderSpecificSettings interface {
	ReplicationClusterProviderSpecificSettings() BaseReplicationClusterProviderSpecificSettingsImpl
}

var _ ReplicationClusterProviderSpecificSettings = BaseReplicationClusterProviderSpecificSettingsImpl{}

type BaseReplicationClusterProviderSpecificSettingsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseReplicationClusterProviderSpecificSettingsImpl) ReplicationClusterProviderSpecificSettings() BaseReplicationClusterProviderSpecificSettingsImpl {
	return s
}

var _ ReplicationClusterProviderSpecificSettings = RawReplicationClusterProviderSpecificSettingsImpl{}

// RawReplicationClusterProviderSpecificSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawReplicationClusterProviderSpecificSettingsImpl struct {
	replicationClusterProviderSpecificSettings BaseReplicationClusterProviderSpecificSettingsImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawReplicationClusterProviderSpecificSettingsImpl) ReplicationClusterProviderSpecificSettings() BaseReplicationClusterProviderSpecificSettingsImpl {
	return s.replicationClusterProviderSpecificSettings
}

func UnmarshalReplicationClusterProviderSpecificSettingsImplementation(input []byte) (ReplicationClusterProviderSpecificSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ReplicationClusterProviderSpecificSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "A2A") {
		var out A2AReplicationProtectionClusterDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into A2AReplicationProtectionClusterDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseReplicationClusterProviderSpecificSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseReplicationClusterProviderSpecificSettingsImpl: %+v", err)
	}

	return RawReplicationClusterProviderSpecificSettingsImpl{
		replicationClusterProviderSpecificSettings: parent,
		Type:   value,
		Values: temp,
	}, nil

}
