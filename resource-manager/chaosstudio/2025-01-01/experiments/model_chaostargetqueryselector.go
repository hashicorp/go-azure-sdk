package experiments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChaosTargetSelector = ChaosTargetQuerySelector{}

type ChaosTargetQuerySelector struct {
	QueryString     string   `json:"queryString"`
	SubscriptionIds []string `json:"subscriptionIds"`

	// Fields inherited from ChaosTargetSelector

	Filter ChaosTargetFilter `json:"filter"`
	Id     string            `json:"id"`
	Type   SelectorType      `json:"type"`
}

func (s ChaosTargetQuerySelector) ChaosTargetSelector() BaseChaosTargetSelectorImpl {
	return BaseChaosTargetSelectorImpl{
		Filter: s.Filter,
		Id:     s.Id,
		Type:   s.Type,
	}
}

var _ json.Marshaler = ChaosTargetQuerySelector{}

func (s ChaosTargetQuerySelector) MarshalJSON() ([]byte, error) {
	type wrapper ChaosTargetQuerySelector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChaosTargetQuerySelector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosTargetQuerySelector: %+v", err)
	}

	decoded["type"] = "Query"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChaosTargetQuerySelector: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChaosTargetQuerySelector{}

func (s *ChaosTargetQuerySelector) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		QueryString     string       `json:"queryString"`
		SubscriptionIds []string     `json:"subscriptionIds"`
		Id              string       `json:"id"`
		Type            SelectorType `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.QueryString = decoded.QueryString
	s.SubscriptionIds = decoded.SubscriptionIds
	s.Id = decoded.Id
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChaosTargetQuerySelector into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["filter"]; ok {
		impl, err := UnmarshalChaosTargetFilterImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Filter' for 'ChaosTargetQuerySelector': %+v", err)
		}
		s.Filter = impl
	}

	return nil
}
