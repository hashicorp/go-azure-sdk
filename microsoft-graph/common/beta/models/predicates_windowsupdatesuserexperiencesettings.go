package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUserExperienceSettingsOperationPredicate struct {
	DaysUntilForcedReboot *int64
	ODataType             *string
}

func (p WindowsUpdatesUserExperienceSettingsOperationPredicate) Matches(input WindowsUpdatesUserExperienceSettings) bool {

	if p.DaysUntilForcedReboot != nil && (input.DaysUntilForcedReboot == nil || *p.DaysUntilForcedReboot != *input.DaysUntilForcedReboot) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
