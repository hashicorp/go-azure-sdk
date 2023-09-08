package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSettingOperationPredicate struct {
	Description       *string
	DisplayName       *string
	DurationInMinutes *string
	ODataType         *string
	Url               *string
}

func (p CustomTrainingSettingOperationPredicate) Matches(input CustomTrainingSetting) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DurationInMinutes != nil && (input.DurationInMinutes == nil || *p.DurationInMinutes != *input.DurationInMinutes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
