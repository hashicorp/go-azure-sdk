package operation

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ValidateOperationRequest = ValidateIaasVMRestoreOperationRequest{}

type ValidateIaasVMRestoreOperationRequest struct {
	RestoreRequest RestoreRequest `json:"restoreRequest"`

	// Fields inherited from ValidateOperationRequest
}

var _ json.Marshaler = ValidateIaasVMRestoreOperationRequest{}

func (s ValidateIaasVMRestoreOperationRequest) MarshalJSON() ([]byte, error) {
	type wrapper ValidateIaasVMRestoreOperationRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateIaasVMRestoreOperationRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateIaasVMRestoreOperationRequest: %+v", err)
	}
	decoded["objectType"] = "ValidateIaasVMRestoreOperationRequest"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateIaasVMRestoreOperationRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ValidateIaasVMRestoreOperationRequest{}

func (s *ValidateIaasVMRestoreOperationRequest) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ValidateIaasVMRestoreOperationRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["restoreRequest"]; ok {
		impl, err := unmarshalRestoreRequestImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'RestoreRequest' for 'ValidateIaasVMRestoreOperationRequest': %+v", err)
		}
		s.RestoreRequest = impl
	}
	return nil
}
