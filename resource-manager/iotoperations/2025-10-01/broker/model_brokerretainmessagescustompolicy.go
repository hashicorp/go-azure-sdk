package broker

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BrokerRetainMessagesPolicy = BrokerRetainMessagesCustomPolicy{}

type BrokerRetainMessagesCustomPolicy struct {
	RetainSettings BrokerRetainMessagesSettings `json:"retainSettings"`

	// Fields inherited from BrokerRetainMessagesPolicy

	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BrokerRetainMessagesCustomPolicy) BrokerRetainMessagesPolicy() BaseBrokerRetainMessagesPolicyImpl {
	return BaseBrokerRetainMessagesPolicyImpl{
		Mode: s.Mode,
	}
}

var _ json.Marshaler = BrokerRetainMessagesCustomPolicy{}

func (s BrokerRetainMessagesCustomPolicy) MarshalJSON() ([]byte, error) {
	type wrapper BrokerRetainMessagesCustomPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrokerRetainMessagesCustomPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerRetainMessagesCustomPolicy: %+v", err)
	}

	decoded["mode"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrokerRetainMessagesCustomPolicy: %+v", err)
	}

	return encoded, nil
}
