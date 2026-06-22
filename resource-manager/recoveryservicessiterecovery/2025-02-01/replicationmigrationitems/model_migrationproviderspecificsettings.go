package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationProviderSpecificSettings interface {
	MigrationProviderSpecificSettings() BaseMigrationProviderSpecificSettingsImpl
}

var _ MigrationProviderSpecificSettings = BaseMigrationProviderSpecificSettingsImpl{}

type BaseMigrationProviderSpecificSettingsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseMigrationProviderSpecificSettingsImpl) MigrationProviderSpecificSettings() BaseMigrationProviderSpecificSettingsImpl {
	return s
}

var _ MigrationProviderSpecificSettings = RawMigrationProviderSpecificSettingsImpl{}

// RawMigrationProviderSpecificSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawMigrationProviderSpecificSettingsImpl struct {
	migrationProviderSpecificSettings BaseMigrationProviderSpecificSettingsImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawMigrationProviderSpecificSettingsImpl) MigrationProviderSpecificSettings() BaseMigrationProviderSpecificSettingsImpl {
	return s.migrationProviderSpecificSettings
}

func (s RawMigrationProviderSpecificSettingsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalMigrationProviderSpecificSettingsImplementation(input []byte) (MigrationProviderSpecificSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrationProviderSpecificSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtMigrationDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtMigrationDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseMigrationProviderSpecificSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMigrationProviderSpecificSettingsImpl: %+v", err)
	}

	return RawMigrationProviderSpecificSettingsImpl{
		migrationProviderSpecificSettings: parent,
		Type:                              value,
		Values:                            temp,
	}, nil

}
