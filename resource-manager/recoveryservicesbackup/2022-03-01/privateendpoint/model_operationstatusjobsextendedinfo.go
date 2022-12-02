package privateendpoint

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OperationStatusExtendedInfo = OperationStatusJobsExtendedInfo{}

type OperationStatusJobsExtendedInfo struct {
	FailedJobsError *map[string]string `json:"failedJobsError,omitempty"`
	JobIds          *[]string          `json:"jobIds,omitempty"`

	// Fields inherited from OperationStatusExtendedInfo
}

var _ json.Marshaler = OperationStatusJobsExtendedInfo{}

func (s OperationStatusJobsExtendedInfo) MarshalJSON() ([]byte, error) {
	type wrapper OperationStatusJobsExtendedInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationStatusJobsExtendedInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationStatusJobsExtendedInfo: %+v", err)
	}
	decoded["objectType"] = "OperationStatusJobsExtendedInfo"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationStatusJobsExtendedInfo: %+v", err)
	}

	return encoded, nil
}
