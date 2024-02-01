package integrationruntime

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SsisObjectMetadataListResponse struct {
	NextLink *string               `json:"nextLink,omitempty"`
	Value    *[]SsisObjectMetadata `json:"value,omitempty"`
}

var _ json.Unmarshaler = &SsisObjectMetadataListResponse{}

func (s *SsisObjectMetadataListResponse) UnmarshalJSON(bytes []byte) error {
	type alias SsisObjectMetadataListResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SsisObjectMetadataListResponse: %+v", err)
	}

	s.NextLink = decoded.NextLink

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SsisObjectMetadataListResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["value"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Value into list []json.RawMessage: %+v", err)
		}

		output := make([]SsisObjectMetadata, 0)
		for i, val := range listTemp {
			impl, err := unmarshalSsisObjectMetadataImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Value' for 'SsisObjectMetadataListResponse': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Value = &output
	}
	return nil
}
