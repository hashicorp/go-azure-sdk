package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceTimelineEventOperationPredicate struct {
	DeviceId      *string
	EventDateTime *string
	EventDetails  *string
	EventName     *string
	EventSource   *string
	Id            *string
	ODataType     *string
}

func (p UserExperienceAnalyticsDeviceTimelineEventOperationPredicate) Matches(input UserExperienceAnalyticsDeviceTimelineEvent) bool {

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.EventDetails != nil && (input.EventDetails == nil || *p.EventDetails != *input.EventDetails) {
		return false
	}

	if p.EventName != nil && (input.EventName == nil || *p.EventName != *input.EventName) {
		return false
	}

	if p.EventSource != nil && (input.EventSource == nil || *p.EventSource != *input.EventSource) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
