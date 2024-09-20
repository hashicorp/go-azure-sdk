package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTimelineResponse struct {
	MetaData *TimelineResultsMetadata `json:"metaData,omitempty"`
	Value    *[]EntityTimelineItem    `json:"value,omitempty"`
}

var _ json.Unmarshaler = &EntityTimelineResponse{}

func (s *EntityTimelineResponse) UnmarshalJSON(bytes []byte) error {
	type alias EntityTimelineResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EntityTimelineResponse: %+v", err)
	}

	s.MetaData = decoded.MetaData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EntityTimelineResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]EntityTimelineItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEntityTimelineItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'EntityTimelineResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
