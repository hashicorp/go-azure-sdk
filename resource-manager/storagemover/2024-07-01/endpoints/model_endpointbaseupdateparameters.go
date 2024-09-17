package endpoints

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointBaseUpdateParameters struct {
	Properties EndpointBaseUpdateProperties `json:"properties"`
}

var _ json.Unmarshaler = &EndpointBaseUpdateParameters{}

func (s *EndpointBaseUpdateParameters) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EndpointBaseUpdateParameters into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalEndpointBaseUpdatePropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'EndpointBaseUpdateParameters': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
