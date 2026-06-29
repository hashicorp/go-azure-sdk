package broker

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerSubscriberQueuePolicy interface {
	BrokerSubscriberQueuePolicy() BaseBrokerSubscriberQueuePolicyImpl
}

var _ BrokerSubscriberQueuePolicy = BaseBrokerSubscriberQueuePolicyImpl{}

type BaseBrokerSubscriberQueuePolicyImpl struct {
	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BaseBrokerSubscriberQueuePolicyImpl) BrokerSubscriberQueuePolicy() BaseBrokerSubscriberQueuePolicyImpl {
	return s
}

var _ BrokerSubscriberQueuePolicy = RawBrokerSubscriberQueuePolicyImpl{}

// RawBrokerSubscriberQueuePolicyImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawBrokerSubscriberQueuePolicyImpl struct {
	brokerSubscriberQueuePolicy BaseBrokerSubscriberQueuePolicyImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawBrokerSubscriberQueuePolicyImpl) BrokerSubscriberQueuePolicy() BaseBrokerSubscriberQueuePolicyImpl {
	return s.brokerSubscriberQueuePolicy
}

func (s RawBrokerSubscriberQueuePolicyImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalBrokerSubscriberQueuePolicyImplementation(input []byte) (BrokerSubscriberQueuePolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerSubscriberQueuePolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Custom") {
		var out BrokerSubscriberQueueCustomPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BrokerSubscriberQueueCustomPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseBrokerSubscriberQueuePolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBrokerSubscriberQueuePolicyImpl: %+v", err)
	}

	return RawBrokerSubscriberQueuePolicyImpl{
		brokerSubscriberQueuePolicy: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
