package broker

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrokerPersistence struct {
	Encryption                *BrokerPersistenceEncryption `json:"encryption,omitempty"`
	MaxSize                   string                       `json:"maxSize"`
	PersistentVolumeClaimSpec *VolumeClaimSpec             `json:"persistentVolumeClaimSpec,omitempty"`
	Retain                    BrokerRetainMessagesPolicy   `json:"retain"`
	StateStore                BrokerStateStorePolicy       `json:"stateStore"`
	SubscriberQueue           BrokerSubscriberQueuePolicy  `json:"subscriberQueue"`
}

var _ json.Unmarshaler = &BrokerPersistence{}

func (s *BrokerPersistence) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Encryption                *BrokerPersistenceEncryption `json:"encryption,omitempty"`
		MaxSize                   string                       `json:"maxSize"`
		PersistentVolumeClaimSpec *VolumeClaimSpec             `json:"persistentVolumeClaimSpec,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Encryption = decoded.Encryption
	s.MaxSize = decoded.MaxSize
	s.PersistentVolumeClaimSpec = decoded.PersistentVolumeClaimSpec

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BrokerPersistence into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["retain"]; ok {
		impl, err := UnmarshalBrokerRetainMessagesPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Retain' for 'BrokerPersistence': %+v", err)
		}
		s.Retain = impl
	}

	if v, ok := temp["stateStore"]; ok {
		impl, err := UnmarshalBrokerStateStorePolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StateStore' for 'BrokerPersistence': %+v", err)
		}
		s.StateStore = impl
	}

	if v, ok := temp["subscriberQueue"]; ok {
		impl, err := UnmarshalBrokerSubscriberQueuePolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SubscriberQueue' for 'BrokerPersistence': %+v", err)
		}
		s.SubscriberQueue = impl
	}

	return nil
}
