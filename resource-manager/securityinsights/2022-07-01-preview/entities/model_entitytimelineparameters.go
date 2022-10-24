package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTimelineParameters struct {
	EndTime        string                `json:"endTime"`
	Kinds          *[]EntityTimelineKind `json:"kinds,omitempty"`
	NumberOfBucket *int64                `json:"numberOfBucket,omitempty"`
	StartTime      string                `json:"startTime"`
}

func (o *EntityTimelineParameters) GetEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityTimelineParameters) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = formatted
}

func (o *EntityTimelineParameters) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityTimelineParameters) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
