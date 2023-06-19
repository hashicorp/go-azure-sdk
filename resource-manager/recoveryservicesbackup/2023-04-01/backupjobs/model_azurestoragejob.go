package backupjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
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

func (o *AzureStorageJob) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureStorageJob) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *AzureStorageJob) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureStorageJob) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
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
