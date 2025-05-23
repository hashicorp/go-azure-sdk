package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateProviderSpecificInput interface {
	MigrateProviderSpecificInput() BaseMigrateProviderSpecificInputImpl
}

var _ MigrateProviderSpecificInput = BaseMigrateProviderSpecificInputImpl{}

type BaseMigrateProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseMigrateProviderSpecificInputImpl) MigrateProviderSpecificInput() BaseMigrateProviderSpecificInputImpl {
	return s
}

var _ MigrateProviderSpecificInput = RawMigrateProviderSpecificInputImpl{}

// RawMigrateProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrateProviderSpecificInputImpl struct {
	migrateProviderSpecificInput BaseMigrateProviderSpecificInputImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawMigrateProviderSpecificInputImpl) MigrateProviderSpecificInput() BaseMigrateProviderSpecificInputImpl {
	return s.migrateProviderSpecificInput
}

func UnmarshalMigrateProviderSpecificInputImplementation(input []byte) (MigrateProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtMigrateInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtMigrateInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseMigrateProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMigrateProviderSpecificInputImpl: %+v", err)
	}

	return RawMigrateProviderSpecificInputImpl{
		migrateProviderSpecificInput: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
