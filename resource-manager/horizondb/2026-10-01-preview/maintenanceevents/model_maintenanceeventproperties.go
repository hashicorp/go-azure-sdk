package maintenanceevents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventProperties struct {
	Deferrable           bool                   `json:"deferrable"`
	DeferralDeadlineUtc  *string                `json:"deferralDeadlineUtc,omitempty"`
	Description          *string                `json:"description,omitempty"`
	EndTimeUtc           string                 `json:"endTimeUtc"`
	EstimatedDowntime    *string                `json:"estimatedDowntime,omitempty"`
	LastUpdatedTimeUtc   *string                `json:"lastUpdatedTimeUtc,omitempty"`
	MaintenanceEventId   *string                `json:"maintenanceEventId,omitempty"`
	MaintenanceType      MaintenanceType        `json:"maintenanceType"`
	OriginalStartTimeUtc string                 `json:"originalStartTimeUtc"`
	RescheduledFromUtc   *string                `json:"rescheduledFromUtc,omitempty"`
	StartTimeUtc         string                 `json:"startTimeUtc"`
	Status               MaintenanceEventStatus `json:"status"`
}

func (o *MaintenanceEventProperties) GetDeferralDeadlineUtcAsTime() (*time.Time, error) {
	if o.DeferralDeadlineUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeferralDeadlineUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetDeferralDeadlineUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeferralDeadlineUtc = &formatted
}

func (o *MaintenanceEventProperties) GetEndTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetEndTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTimeUtc = formatted
}

func (o *MaintenanceEventProperties) GetLastUpdatedTimeUtcAsTime() (*time.Time, error) {
	if o.LastUpdatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetLastUpdatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedTimeUtc = &formatted
}

func (o *MaintenanceEventProperties) GetOriginalStartTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.OriginalStartTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetOriginalStartTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.OriginalStartTimeUtc = formatted
}

func (o *MaintenanceEventProperties) GetRescheduledFromUtcAsTime() (*time.Time, error) {
	if o.RescheduledFromUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RescheduledFromUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetRescheduledFromUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RescheduledFromUtc = &formatted
}

func (o *MaintenanceEventProperties) GetStartTimeUtcAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *MaintenanceEventProperties) SetStartTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTimeUtc = formatted
}
