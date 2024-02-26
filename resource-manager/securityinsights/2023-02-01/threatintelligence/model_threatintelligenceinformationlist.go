package threatintelligence

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatIntelligenceInformationList struct {
	NextLink *string                         `json:"nextLink,omitempty"`
	Value    []ThreatIntelligenceInformation `json:"value"`
}

var _ json.Unmarshaler = &ThreatIntelligenceInformationList{}

func (s *ThreatIntelligenceInformationList) UnmarshalJSON(bytes []byte) error {
	type alias ThreatIntelligenceInformationList
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ThreatIntelligenceInformationList: %+v", err)
	}

	s.NextLink = decoded.NextLink

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ThreatIntelligenceInformationList into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]ThreatIntelligenceInformation, 0)
		for i, val := range listTemp {
			impl, err := unmarshalThreatIntelligenceInformationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'ThreatIntelligenceInformationList': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = output
	}
	return nil
}
