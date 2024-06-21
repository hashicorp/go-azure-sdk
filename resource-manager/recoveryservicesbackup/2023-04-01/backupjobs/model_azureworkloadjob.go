package backupjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = AzureWorkloadJob{}

type AzureWorkloadJob struct {
	ActionsInfo  *[]JobSupportedAction         `json:"actionsInfo,omitempty"`
	Duration     *string                       `json:"duration,omitempty"`
	ErrorDetails *[]AzureWorkloadErrorInfo     `json:"errorDetails,omitempty"`
	ExtendedInfo *AzureWorkloadJobExtendedInfo `json:"extendedInfo,omitempty"`
	WorkloadType *string                       `json:"workloadType,omitempty"`

	// Fields inherited from Job
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

var _ json.Marshaler = AzureWorkloadJob{}

func (s AzureWorkloadJob) MarshalJSON() ([]byte, error) {
	type wrapper AzureWorkloadJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureWorkloadJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureWorkloadJob: %+v", err)
	}
	decoded["jobType"] = "AzureWorkloadJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureWorkloadJob: %+v", err)
	}

	return encoded, nil
}
