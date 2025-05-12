package replicationmigrationitems

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResyncProviderSpecificInput interface {
	ResyncProviderSpecificInput() BaseResyncProviderSpecificInputImpl
}

var _ ResyncProviderSpecificInput = BaseResyncProviderSpecificInputImpl{}

type BaseResyncProviderSpecificInputImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseResyncProviderSpecificInputImpl) ResyncProviderSpecificInput() BaseResyncProviderSpecificInputImpl {
	return s
}

var _ ResyncProviderSpecificInput = RawResyncProviderSpecificInputImpl{}

// RawResyncProviderSpecificInputImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawResyncProviderSpecificInputImpl struct {
	resyncProviderSpecificInput BaseResyncProviderSpecificInputImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawResyncProviderSpecificInputImpl) ResyncProviderSpecificInput() BaseResyncProviderSpecificInputImpl {
	return s.resyncProviderSpecificInput
}

func UnmarshalResyncProviderSpecificInputImplementation(input []byte) (ResyncProviderSpecificInput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ResyncProviderSpecificInput into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "VMwareCbt") {
		var out VMwareCbtResyncInput
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMwareCbtResyncInput: %+v", err)
		}
		return out, nil
	}

	var parent BaseResyncProviderSpecificInputImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseResyncProviderSpecificInputImpl: %+v", err)
	}

	return RawResyncProviderSpecificInputImpl{
		resyncProviderSpecificInput: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
