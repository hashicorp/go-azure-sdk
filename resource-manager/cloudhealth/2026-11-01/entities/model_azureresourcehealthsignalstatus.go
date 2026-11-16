package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureResourceHealthSignalStatus struct {
	AdditionalContext        *string                          `json:"additionalContext,omitempty"`
	AvailabilityReportedTime *string                          `json:"availabilityReportedTime,omitempty"`
	AvailabilityState        *ResourceHealthAvailabilityState `json:"availabilityState,omitempty"`
	Category                 *ResourceHealthCategory          `json:"category,omitempty"`
	DetailedStatus           *string                          `json:"detailedStatus,omitempty"`
	Error                    *string                          `json:"error,omitempty"`
	HealthState              *HealthState                     `json:"healthState,omitempty"`
	ReasonChronicity         *ResourceHealthReasonChronicity  `json:"reasonChronicity,omitempty"`
	ReasonType               *ResourceHealthReasonType        `json:"reasonType,omitempty"`
	ReportedAt               *string                          `json:"reportedAt,omitempty"`
	Summary                  *string                          `json:"summary,omitempty"`
	Value                    *float64                         `json:"value,omitempty"`
}

func (o *AzureResourceHealthSignalStatus) GetAvailabilityReportedTimeAsTime() (*time.Time, error) {
	if o.AvailabilityReportedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AvailabilityReportedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureResourceHealthSignalStatus) SetAvailabilityReportedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AvailabilityReportedTime = &formatted
}

func (o *AzureResourceHealthSignalStatus) GetReportedAtAsTime() (*time.Time, error) {
	if o.ReportedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReportedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureResourceHealthSignalStatus) SetReportedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReportedAt = &formatted
}
