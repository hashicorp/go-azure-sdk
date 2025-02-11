package experiments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChaosExperimentBranch struct {
	Actions []ChaosExperimentAction `json:"actions"`
	Name    string                  `json:"name"`
}

var _ json.Unmarshaler = &ChaosExperimentBranch{}

func (s *ChaosExperimentBranch) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Name = decoded.Name

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChaosExperimentBranch into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]ChaosExperimentAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalChaosExperimentActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'ChaosExperimentBranch': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = output
	}

	return nil
}
