package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMigrationItemProviderSpecificInput interface {
	UpdateMigrationItemProviderSpecificInput() BaseUpdateMigrationItemProviderSpecificInputImpl
}

var _ UpdateMigrationItemProviderSpecificInput = BaseUpdateMigrationItemProviderSpecificInputImpl{}

type BaseUpdateMigrationItemProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseUpdateMigrationItemProviderSpecificInputImpl) UpdateMigrationItemProviderSpecificInput() BaseUpdateMigrationItemProviderSpecificInputImpl {
	return s
}

var _ UpdateMigrationItemProviderSpecificInput = RawUpdateMigrationItemProviderSpecificInputImpl{}

// RawUpdateMigrationItemProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawUpdateMigrationItemProviderSpecificInputImpl struct {
	updateMigrationItemProviderSpecificInput BaseUpdateMigrationItemProviderSpecificInputImpl
	Type                                     string
	Values                                   map[string]interface{}
}

func (s RawUpdateMigrationItemProviderSpecificInputImpl) UpdateMigrationItemProviderSpecificInput() BaseUpdateMigrationItemProviderSpecificInputImpl {
	return s.updateMigrationItemProviderSpecificInput
}

func (s RawUpdateMigrationItemProviderSpecificInputImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalUpdateMigrationItemProviderSpecificInputImplementation(input []byte) (UpdateMigrationItemProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UpdateMigrationItemProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtUpdateMigrationItemInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtUpdateMigrationItemInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseUpdateMigrationItemProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUpdateMigrationItemProviderSpecificInputImpl: %+v", err)
	}

	return RawUpdateMigrationItemProviderSpecificInputImpl{
		updateMigrationItemProviderSpecificInput: parent,
		Type:                                     value,
		Values:                                   temp,
	}, nil

}
