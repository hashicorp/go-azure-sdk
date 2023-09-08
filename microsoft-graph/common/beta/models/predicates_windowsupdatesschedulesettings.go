package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesScheduleSettingsOperationPredicate struct {
	ODataType     *string
	StartDateTime *string
}

func (p WindowsUpdatesScheduleSettingsOperationPredicate) Matches(input WindowsUpdatesScheduleSettings) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
