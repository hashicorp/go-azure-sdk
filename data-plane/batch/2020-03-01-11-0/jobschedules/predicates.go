package jobschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudJobScheduleOperationPredicate struct {
	CreationTime                *string
	DisplayName                 *string
	ETag                        *string
	Id                          *string
	LastModified                *string
	PreviousStateTransitionTime *string
	StateTransitionTime         *string
	Url                         *string
}

func (p CloudJobScheduleOperationPredicate) Matches(input CloudJobSchedule) bool {

	if p.CreationTime != nil && (input.CreationTime == nil || *p.CreationTime != *input.CreationTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ETag != nil && (input.ETag == nil || *p.ETag != *input.ETag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModified != nil && (input.LastModified == nil || *p.LastModified != *input.LastModified) {
		return false
	}

	if p.PreviousStateTransitionTime != nil && (input.PreviousStateTransitionTime == nil || *p.PreviousStateTransitionTime != *input.PreviousStateTransitionTime) {
		return false
	}

	if p.StateTransitionTime != nil && (input.StateTransitionTime == nil || *p.StateTransitionTime != *input.StateTransitionTime) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
