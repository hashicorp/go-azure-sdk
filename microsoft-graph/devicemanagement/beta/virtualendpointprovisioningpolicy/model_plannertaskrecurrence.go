package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRecurrence struct {
	NextInSeriesTaskId      *string                    `json:"nextInSeriesTaskId,omitempty"`
	ODataType               *string                    `json:"@odata.type,omitempty"`
	OccurrenceId            *int64                     `json:"occurrenceId,omitempty"`
	PreviousInSeriesTaskId  *string                    `json:"previousInSeriesTaskId,omitempty"`
	RecurrenceStartDateTime *string                    `json:"recurrenceStartDateTime,omitempty"`
	Schedule                *PlannerRecurrenceSchedule `json:"schedule,omitempty"`
	SeriesId                *string                    `json:"seriesId,omitempty"`
}
