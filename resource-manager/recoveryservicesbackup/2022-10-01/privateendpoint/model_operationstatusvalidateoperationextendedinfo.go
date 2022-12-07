package privateendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OperationStatusExtendedInfo = OperationStatusValidateOperationExtendedInfo{}

type OperationStatusValidateOperationExtendedInfo struct {
	ValidateOperationResponse *ValidateOperationResponse `json:"validateOperationResponse,omitempty"`

	// Fields inherited from OperationStatusExtendedInfo
}

var _ json.Marshaler = OperationStatusValidateOperationExtendedInfo{}

func (s OperationStatusValidateOperationExtendedInfo) MarshalJSON() ([]byte, error) {
	type wrapper OperationStatusValidateOperationExtendedInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationStatusValidateOperationExtendedInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusValidateOperationExtendedInfo: %+v", err)
	}
	decoded["objectType"] = "OperationStatusValidateOperationExtendedInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationStatusValidateOperationExtendedInfo: %+v", err)
	}

	return encoded, nil
}
