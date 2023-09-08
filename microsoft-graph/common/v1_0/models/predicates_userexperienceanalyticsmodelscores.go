package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsModelScoresOperationPredicate struct {
	Id               *string
	Manufacturer     *string
	Model            *string
	ModelDeviceCount *int64
	ODataType        *string
}

func (p UserExperienceAnalyticsModelScoresOperationPredicate) Matches(input UserExperienceAnalyticsModelScores) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Model != nil && (input.Model == nil || *p.Model != *input.Model) {
		return false
	}

	if p.ModelDeviceCount != nil && (input.ModelDeviceCount == nil || *p.ModelDeviceCount != *input.ModelDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
