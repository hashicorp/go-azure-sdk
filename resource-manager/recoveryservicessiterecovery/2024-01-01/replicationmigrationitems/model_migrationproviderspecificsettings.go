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

// RawMigrationProviderSpecificSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrationProviderSpecificSettingsImpl struct {
	migrationProviderSpecificSettings BaseMigrationProviderSpecificSettingsImpl
	Type                              string
	Values                            map[string]interface{}
}

func (s RawMigrationProviderSpecificSettingsImpl) MigrationProviderSpecificSettings() BaseMigrationProviderSpecificSettingsImpl {
	return s.migrationProviderSpecificSettings
}

func UnmarshalMigrationProviderSpecificSettingsImplementation(input []byte) (MigrationProviderSpecificSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrationProviderSpecificSettings into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
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
