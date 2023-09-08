package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDailyScheduleOperationPredicate struct {
	Interval  *int64
	ODataType *string
	Time      *string
	UseUtc    *bool
}

func (p DeviceHealthScriptDailyScheduleOperationPredicate) Matches(input DeviceHealthScriptDailySchedule) bool {

	if p.Interval != nil && (input.Interval == nil || *p.Interval != *input.Interval) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Time != nil && (input.Time == nil || *p.Time != *input.Time) {
		return false
	}

	if p.UseUtc != nil && (input.UseUtc == nil || *p.UseUtc != *input.UseUtc) {
		return false
	}

	return true
}
