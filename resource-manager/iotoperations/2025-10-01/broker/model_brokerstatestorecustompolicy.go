package broker

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BrokerStateStorePolicy = BrokerStateStoreCustomPolicy{}

type BrokerStateStoreCustomPolicy struct {
	StateStoreSettings BrokerStateStorePolicySettings `json:"stateStoreSettings"`

	// Fields inherited from BrokerStateStorePolicy

	Mode BrokerPersistencePolicyMode `json:"mode"`
}

func (s BrokerStateStoreCustomPolicy) BrokerStateStorePolicy() BaseBrokerStateStorePolicyImpl {
	return BaseBrokerStateStorePolicyImpl{
		Mode: s.Mode,
	}
}

var _ json.Marshaler = BrokerStateStoreCustomPolicy{}

func (s BrokerStateStoreCustomPolicy) MarshalJSON() ([]byte, error) {
	type wrapper BrokerStateStoreCustomPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BrokerStateStoreCustomPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BrokerStateStoreCustomPolicy: %+v", err)
	}

	decoded["mode"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BrokerStateStoreCustomPolicy: %+v", err)
	}

	return encoded, nil
}
