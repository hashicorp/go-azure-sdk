package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsImpactingProcessOperationPredicate struct {
	Category    *string
	Description *string
	DeviceId    *string
	Id          *string
	ODataType   *string
	ProcessName *string
	Publisher   *string
}

func (p UserExperienceAnalyticsImpactingProcessOperationPredicate) Matches(input UserExperienceAnalyticsImpactingProcess) bool {

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProcessName != nil && (input.ProcessName == nil || *p.ProcessName != *input.ProcessName) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	return true
}
