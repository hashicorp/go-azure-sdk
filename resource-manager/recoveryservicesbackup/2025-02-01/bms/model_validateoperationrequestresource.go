package bms

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationRequestResource struct {
	Id         string                   `json:"id"`
	Properties ValidateOperationRequest `json:"properties"`
}

var _ json.Unmarshaler = &ValidateOperationRequestResource{}

func (s *ValidateOperationRequestResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id string `json:"id"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ValidateOperationRequestResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalValidateOperationRequestImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'ValidateOperationRequestResource': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
