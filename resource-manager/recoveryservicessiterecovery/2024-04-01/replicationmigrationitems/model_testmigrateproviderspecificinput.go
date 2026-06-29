package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestMigrateProviderSpecificInput interface {
	TestMigrateProviderSpecificInput() BaseTestMigrateProviderSpecificInputImpl
}

var _ TestMigrateProviderSpecificInput = BaseTestMigrateProviderSpecificInputImpl{}

type BaseTestMigrateProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseTestMigrateProviderSpecificInputImpl) TestMigrateProviderSpecificInput() BaseTestMigrateProviderSpecificInputImpl {
	return s
}

var _ TestMigrateProviderSpecificInput = RawTestMigrateProviderSpecificInputImpl{}

// RawTestMigrateProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawTestMigrateProviderSpecificInputImpl struct {
	testMigrateProviderSpecificInput BaseTestMigrateProviderSpecificInputImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawTestMigrateProviderSpecificInputImpl) TestMigrateProviderSpecificInput() BaseTestMigrateProviderSpecificInputImpl {
	return s.testMigrateProviderSpecificInput
}

func (s RawTestMigrateProviderSpecificInputImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalTestMigrateProviderSpecificInputImplementation(input []byte) (TestMigrateProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TestMigrateProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtTestMigrateInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtTestMigrateInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseTestMigrateProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTestMigrateProviderSpecificInputImpl: %+v", err)
	}

	return RawTestMigrateProviderSpecificInputImpl{
		testMigrateProviderSpecificInput: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
