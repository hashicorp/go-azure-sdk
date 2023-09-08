package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRecurrenceOperationPredicate struct {
	NextInSeriesTaskId      *string
	ODataType               *string
	OccurrenceId            *int64
	PreviousInSeriesTaskId  *string
	RecurrenceStartDateTime *string
	SeriesId                *string
}

func (p PlannerTaskRecurrenceOperationPredicate) Matches(input PlannerTaskRecurrence) bool {

	if p.NextInSeriesTaskId != nil && (input.NextInSeriesTaskId == nil || *p.NextInSeriesTaskId != *input.NextInSeriesTaskId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OccurrenceId != nil && (input.OccurrenceId == nil || *p.OccurrenceId != *input.OccurrenceId) {
		return false
	}

	if p.PreviousInSeriesTaskId != nil && (input.PreviousInSeriesTaskId == nil || *p.PreviousInSeriesTaskId != *input.PreviousInSeriesTaskId) {
		return false
	}

	if p.RecurrenceStartDateTime != nil && (input.RecurrenceStartDateTime == nil || *p.RecurrenceStartDateTime != *input.RecurrenceStartDateTime) {
		return false
	}

	if p.SeriesId != nil && (input.SeriesId == nil || *p.SeriesId != *input.SeriesId) {
		return false
	}

	return true
}
