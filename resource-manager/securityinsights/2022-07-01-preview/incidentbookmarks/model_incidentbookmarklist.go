package incidentbookmarks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentBookmarkList struct {
	Value []Entity `json:"value"`
}

var _ json.Unmarshaler = &IncidentBookmarkList{}

func (s *IncidentBookmarkList) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IncidentBookmarkList into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]Entity, 0)
		for i, val := range listTemp {
			impl, err := unmarshalEntityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'IncidentBookmarkList': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = output
	}
	return nil
}
