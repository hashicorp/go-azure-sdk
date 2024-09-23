package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableMigrationProviderSpecificInput interface {
	EnableMigrationProviderSpecificInput() BaseEnableMigrationProviderSpecificInputImpl
}

var _ EnableMigrationProviderSpecificInput = BaseEnableMigrationProviderSpecificInputImpl{}

type BaseEnableMigrationProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseEnableMigrationProviderSpecificInputImpl) EnableMigrationProviderSpecificInput() BaseEnableMigrationProviderSpecificInputImpl {
	return s
}

var _ EnableMigrationProviderSpecificInput = RawEnableMigrationProviderSpecificInputImpl{}

// RawEnableMigrationProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEnableMigrationProviderSpecificInputImpl struct {
	enableMigrationProviderSpecificInput BaseEnableMigrationProviderSpecificInputImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawEnableMigrationProviderSpecificInputImpl) EnableMigrationProviderSpecificInput() BaseEnableMigrationProviderSpecificInputImpl {
	return s.enableMigrationProviderSpecificInput
}

func UnmarshalEnableMigrationProviderSpecificInputImplementation(input []byte) (EnableMigrationProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EnableMigrationProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtEnableMigrationInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtEnableMigrationInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseEnableMigrationProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEnableMigrationProviderSpecificInputImpl: %+v", err)
	}

	return RawEnableMigrationProviderSpecificInputImpl{
		enableMigrationProviderSpecificInput: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
