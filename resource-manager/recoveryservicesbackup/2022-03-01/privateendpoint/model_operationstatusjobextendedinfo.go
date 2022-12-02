package privateendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OperationStatusExtendedInfo = OperationStatusJobExtendedInfo{}

type OperationStatusJobExtendedInfo struct {
	JobId *string `json:"jobId,omitempty"`

	// Fields inherited from OperationStatusExtendedInfo
}

var _ json.Marshaler = OperationStatusJobExtendedInfo{}

func (s OperationStatusJobExtendedInfo) MarshalJSON() ([]byte, error) {
	type wrapper OperationStatusJobExtendedInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationStatusJobExtendedInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusJobExtendedInfo: %+v", err)
	}
	decoded["objectType"] = "OperationStatusJobExtendedInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationStatusJobExtendedInfo: %+v", err)
	}

	return encoded, nil
}
