package broker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerStateStorePolicy interface {
	BrokerStateStorePolicy() BaseBrokerStateStorePolicyImpl
}

var _ BrokerStateStorePolicy = BaseBrokerStateStorePolicyImpl{}

type BaseBrokerStateStorePolicyImpl struct {
	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BaseBrokerStateStorePolicyImpl) BrokerStateStorePolicy() BaseBrokerStateStorePolicyImpl {
	return s
}

var _ BrokerStateStorePolicy = RawBrokerStateStorePolicyImpl{}

// RawBrokerStateStorePolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBrokerStateStorePolicyImpl struct {
	brokerStateStorePolicy BaseBrokerStateStorePolicyImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawBrokerStateStorePolicyImpl) BrokerStateStorePolicy() BaseBrokerStateStorePolicyImpl {
	return s.brokerStateStorePolicy
}

func UnmarshalBrokerStateStorePolicyImplementation(input []byte) (BrokerStateStorePolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerStateStorePolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Custom") {
		var out BrokerStateStoreCustomPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrokerStateStoreCustomPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseBrokerStateStorePolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBrokerStateStorePolicyImpl: %+v", err)
	}

	return RawBrokerStateStorePolicyImpl{
		brokerStateStorePolicy: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
