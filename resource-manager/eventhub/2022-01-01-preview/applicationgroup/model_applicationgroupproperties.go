package applicationgroup

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGroupProperties struct {
	ClientAppGroupIdentifier string                    `json:"clientAppGroupIdentifier"`
	IsEnabled                *bool                     `json:"isEnabled,omitempty"`
	Policies                 *[]ApplicationGroupPolicy `json:"policies,omitempty"`
}

var _ json.Unmarshaler = &ApplicationGroupProperties{}

func (s *ApplicationGroupProperties) UnmarshalJSON(bytes []byte) error {
	type alias ApplicationGroupProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ApplicationGroupProperties: %+v", err)
	}

	s.ClientAppGroupIdentifier = decoded.ClientAppGroupIdentifier
	s.IsEnabled = decoded.IsEnabled

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ApplicationGroupProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Policies into list []json.RawMessage: %+v", err)
		}

		output := make([]ApplicationGroupPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalApplicationGroupPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Policies' for 'ApplicationGroupProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Policies = &output
	}
	return nil
}
