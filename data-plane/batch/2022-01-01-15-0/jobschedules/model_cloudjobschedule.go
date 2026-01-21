package jobschedules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudJobSchedule struct {
	CreationTime                *string                          `json:"creationTime,omitempty"`
	DisplayName                 *string                          `json:"displayName,omitempty"`
	ETag                        *string                          `json:"eTag,omitempty"`
	ExecutionInfo               *JobScheduleExecutionInformation `json:"executionInfo,omitempty"`
	Id                          *string                          `json:"id,omitempty"`
	JobSpecification            *JobSpecification                `json:"jobSpecification,omitempty"`
	LastModified                *string                          `json:"lastModified,omitempty"`
	Metadata                    *[]MetadataItem                  `json:"metadata,omitempty"`
	PreviousState               *JobScheduleState                `json:"previousState,omitempty"`
	PreviousStateTransitionTime *string                          `json:"previousStateTransitionTime,omitempty"`
	Schedule                    *Schedule                        `json:"schedule,omitempty"`
	State                       *JobScheduleState                `json:"state,omitempty"`
	StateTransitionTime         *string                          `json:"stateTransitionTime,omitempty"`
	Stats                       *JobScheduleStatistics           `json:"stats,omitempty"`
	Url                         *string                          `json:"url,omitempty"`
}

func (o *CloudJobSchedule) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudJobSchedule) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *CloudJobSchedule) GetLastModifiedAsTime() (*time.Time, error) {
	if o.LastModified == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModified, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudJobSchedule) SetLastModifiedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModified = &formatted
}

func (o *CloudJobSchedule) GetPreviousStateTransitionTimeAsTime() (*time.Time, error) {
	if o.PreviousStateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PreviousStateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudJobSchedule) SetPreviousStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PreviousStateTransitionTime = &formatted
}

func (o *CloudJobSchedule) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CloudJobSchedule) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
