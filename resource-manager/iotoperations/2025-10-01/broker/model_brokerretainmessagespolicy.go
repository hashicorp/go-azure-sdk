package broker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerRetainMessagesPolicy interface {
	BrokerRetainMessagesPolicy() BaseBrokerRetainMessagesPolicyImpl
}

var _ BrokerRetainMessagesPolicy = BaseBrokerRetainMessagesPolicyImpl{}

type BaseBrokerRetainMessagesPolicyImpl struct {
	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BaseBrokerRetainMessagesPolicyImpl) BrokerRetainMessagesPolicy() BaseBrokerRetainMessagesPolicyImpl {
	return s
}

var _ BrokerRetainMessagesPolicy = RawBrokerRetainMessagesPolicyImpl{}

// RawBrokerRetainMessagesPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBrokerRetainMessagesPolicyImpl struct {
	brokerRetainMessagesPolicy BaseBrokerRetainMessagesPolicyImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawBrokerRetainMessagesPolicyImpl) BrokerRetainMessagesPolicy() BaseBrokerRetainMessagesPolicyImpl {
	return s.brokerRetainMessagesPolicy
}

func UnmarshalBrokerRetainMessagesPolicyImplementation(input []byte) (BrokerRetainMessagesPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerRetainMessagesPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Custom") {
		var out BrokerRetainMessagesCustomPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrokerRetainMessagesCustomPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseBrokerRetainMessagesPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBrokerRetainMessagesPolicyImpl: %+v", err)
	}

	return RawBrokerRetainMessagesPolicyImpl{
		brokerRetainMessagesPolicy: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
