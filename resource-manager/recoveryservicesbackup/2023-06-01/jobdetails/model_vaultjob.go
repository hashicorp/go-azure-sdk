package jobdetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = VaultJob{}

type VaultJob struct {
	ActionsInfo  *[]JobSupportedAction `json:"actionsInfo,omitempty"`
	Duration     *string               `json:"duration,omitempty"`
	ErrorDetails *[]VaultJobErrorInfo  `json:"errorDetails,omitempty"`
	ExtendedInfo *VaultJobExtendedInfo `json:"extendedInfo,omitempty"`

	// Fields inherited from Job
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

var _ json.Marshaler = VaultJob{}

func (s VaultJob) MarshalJSON() ([]byte, error) {
	type wrapper VaultJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VaultJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VaultJob: %+v", err)
	}
	decoded["jobType"] = "VaultJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VaultJob: %+v", err)
	}

	return encoded, nil
}
