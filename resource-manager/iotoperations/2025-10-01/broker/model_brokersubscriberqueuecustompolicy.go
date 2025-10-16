package broker

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BrokerSubscriberQueuePolicy = BrokerSubscriberQueueCustomPolicy{}

type BrokerSubscriberQueueCustomPolicy struct {
	SubscriberQueueSettings BrokerSubscriberQueueCustomPolicySettings `json:"subscriberQueueSettings"`

	// Fields inherited from BrokerSubscriberQueuePolicy

	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BrokerSubscriberQueueCustomPolicy) BrokerSubscriberQueuePolicy() BaseBrokerSubscriberQueuePolicyImpl {
	return BaseBrokerSubscriberQueuePolicyImpl{
		Mode: s.Mode,
	}
}

var _ json.Marshaler = BrokerSubscriberQueueCustomPolicy{}

func (s BrokerSubscriberQueueCustomPolicy) MarshalJSON() ([]byte, error) {
	type wrapper BrokerSubscriberQueueCustomPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrokerSubscriberQueueCustomPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerSubscriberQueueCustomPolicy: %+v", err)
	}

	decoded["mode"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrokerSubscriberQueueCustomPolicy: %+v", err)
	}

	return encoded, nil
}
