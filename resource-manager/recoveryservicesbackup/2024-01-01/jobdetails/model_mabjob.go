package jobdetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = MabJob{}

type MabJob struct {
	ActionsInfo   *[]JobSupportedAction `json:"actionsInfo,omitempty"`
	Duration      *string               `json:"duration,omitempty"`
	ErrorDetails  *[]MabErrorInfo       `json:"errorDetails,omitempty"`
	ExtendedInfo  *MabJobExtendedInfo   `json:"extendedInfo,omitempty"`
	MabServerName *string               `json:"mabServerName,omitempty"`
	MabServerType *MabServerType        `json:"mabServerType,omitempty"`
	WorkloadType  *WorkloadType         `json:"workloadType,omitempty"`

	// Fields inherited from Job
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

var _ json.Marshaler = MabJob{}

func (s MabJob) MarshalJSON() ([]byte, error) {
	type wrapper MabJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MabJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MabJob: %+v", err)
	}
	decoded["jobType"] = "MabJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MabJob: %+v", err)
	}

	return encoded, nil
}
