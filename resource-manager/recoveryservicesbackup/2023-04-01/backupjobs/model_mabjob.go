package backupjobs

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
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
	JobType              string                `json:"jobType"`
	Operation            *string               `json:"operation,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
}

func (s MabJob) Job() BaseJobImpl {
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

func (o *MabJob) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MabJob) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *MabJob) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MabJob) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MabJob: %+v", err)
	}

	decoded["jobType"] = "MabJob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MabJob: %+v", err)
	}

	return encoded, nil
}
