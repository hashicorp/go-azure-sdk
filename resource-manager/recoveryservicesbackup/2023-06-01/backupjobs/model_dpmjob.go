package backupjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Job = DpmJob{}

type DpmJob struct {
	ActionsInfo   *[]JobSupportedAction `json:"actionsInfo,omitempty"`
	ContainerName *string               `json:"containerName,omitempty"`
	ContainerType *string               `json:"containerType,omitempty"`
	DpmServerName *string               `json:"dpmServerName,omitempty"`
	Duration      *string               `json:"duration,omitempty"`
	ErrorDetails  *[]DpmErrorInfo       `json:"errorDetails,omitempty"`
	ExtendedInfo  *DpmJobExtendedInfo   `json:"extendedInfo,omitempty"`
	WorkloadType  *string               `json:"workloadType,omitempty"`

	// Fields inherited from Job

	ActivityId           *string               `json:"activityId,omitempty"`
	BackupManagementType *BackupManagementType `json:"backupManagementType,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	EntityFriendlyName   *string               `json:"entityFriendlyName,omitempty"`
	JobType              string                `json:"jobType"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

func (s DpmJob) Job() BaseJobImpl {
	return BaseJobImpl{
		ActivityId:           s.ActivityId,
		BackupManagementType: s.BackupManagementType,
		EndTime:              s.EndTime,
		EntityFriendlyName:   s.EntityFriendlyName,
		JobType:              s.JobType,
		Operation:            s.Operation,
		StartTime:            s.StartTime,
		Status:               s.Status,
	}
}

func (o *DpmJob) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DpmJob) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *DpmJob) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DpmJob) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

var _ json.Marshaler = DpmJob{}

func (s DpmJob) MarshalJSON() ([]byte, error) {
	type wrapper DpmJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DpmJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DpmJob: %+v", err)
	}

	decoded["jobType"] = "DpmJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DpmJob: %+v", err)
	}

	return encoded, nil
}
