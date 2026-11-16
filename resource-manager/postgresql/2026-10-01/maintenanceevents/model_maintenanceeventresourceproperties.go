package maintenanceevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventResourceProperties struct {
	Deferrable         bool                   `json:"deferrable"`
	DeferralDeadline   *string                `json:"deferralDeadline,omitempty"`
	Description        *string                `json:"description,omitempty"`
	EndTime            string                 `json:"endTime"`
	EstimatedDowntime  *string                `json:"estimatedDowntime,omitempty"`
	LastUpdatedTime    *string                `json:"lastUpdatedTime,omitempty"`
	MaintenanceEventId *string                `json:"maintenanceEventId,omitempty"`
	MaintenanceType    MaintenanceType        `json:"maintenanceType"`
	OriginalStartTime  string                 `json:"originalStartTime"`
	RescheduledFrom    *string                `json:"rescheduledFrom,omitempty"`
	StartTime          string                 `json:"startTime"`
	Status             MaintenanceEventStatus `json:"status"`
}

func (o *MaintenanceEventResourceProperties) GetDeferralDeadlineAsTime() (*time.Time, error) {
	if o.DeferralDeadline == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeferralDeadline, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetDeferralDeadlineAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeferralDeadline = &formatted
}

func (o *MaintenanceEventResourceProperties) GetEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = formatted
}

func (o *MaintenanceEventResourceProperties) GetLastUpdatedTimeAsTime() (*time.Time, error) {
	if o.LastUpdatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetLastUpdatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedTime = &formatted
}

func (o *MaintenanceEventResourceProperties) GetOriginalStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.OriginalStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetOriginalStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.OriginalStartTime = formatted
}

func (o *MaintenanceEventResourceProperties) GetRescheduledFromAsTime() (*time.Time, error) {
	if o.RescheduledFrom == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RescheduledFrom, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetRescheduledFromAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RescheduledFrom = &formatted
}

func (o *MaintenanceEventResourceProperties) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventResourceProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
