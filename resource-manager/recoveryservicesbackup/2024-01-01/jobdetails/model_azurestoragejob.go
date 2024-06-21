package jobdetails

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = AzureStorageJob{}

type AzureStorageJob struct {
	ActionsInfo           *[]JobSupportedAction        `json:"actionsInfo,omitempty"`
	Duration              *string                      `json:"duration,omitempty"`
	ErrorDetails          *[]AzureStorageErrorInfo     `json:"errorDetails,omitempty"`
	ExtendedInfo          *AzureStorageJobExtendedInfo `json:"extendedInfo,omitempty"`
	IsUserTriggered       *bool                        `json:"isUserTriggered,omitempty"`
	StorageAccountName    *string                      `json:"storageAccountName,omitempty"`
	StorageAccountVersion *string                      `json:"storageAccountVersion,omitempty"`

	// Fields inherited from Job
	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

var _ json.Marshaler = AzureStorageJob{}

func (s AzureStorageJob) MarshalJSON() ([]byte, error) {
	type wrapper AzureStorageJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureStorageJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureStorageJob: %+v", err)
	}
	decoded["jobType"] = "AzureStorageJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureStorageJob: %+v", err)
	}

	return encoded, nil
}
