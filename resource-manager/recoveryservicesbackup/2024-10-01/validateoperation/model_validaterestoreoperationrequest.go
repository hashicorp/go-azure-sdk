package validateoperation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ValidateOperationRequest = ValidateRestoreOperationRequest{}

type ValidateRestoreOperationRequest struct {
	RestoreRequest RestoreRequest `json:"restoreRequest"`

	// Fields inherited from ValidateOperationRequest

	ObjectType string `json:"objectType"`
}

func (s ValidateRestoreOperationRequest) ValidateOperationRequest() BaseValidateOperationRequestImpl {
	return BaseValidateOperationRequestImpl{
		ObjectType: s.ObjectType,
	}
}

var _ json.Marshaler = ValidateRestoreOperationRequest{}

func (s ValidateRestoreOperationRequest) MarshalJSON() ([]byte, error) {
	type wrapper ValidateRestoreOperationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateRestoreOperationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateRestoreOperationRequest: %+v", err)
	}

	decoded["objectType"] = "ValidateRestoreOperationRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateRestoreOperationRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ValidateRestoreOperationRequest{}

func (s *ValidateRestoreOperationRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ObjectType string `json:"objectType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ObjectType = decoded.ObjectType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ValidateRestoreOperationRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["restoreRequest"]; ok {
		impl, err := UnmarshalRestoreRequestImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RestoreRequest' for 'ValidateRestoreOperationRequest': %+v", err)
		}
		s.RestoreRequest = impl
	}

	return nil
}
