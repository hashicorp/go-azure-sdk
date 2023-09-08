package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpeditedWindowsQualityUpdateSettingsOperationPredicate struct {
	DaysUntilForcedReboot *int64
	ODataType             *string
	QualityUpdateRelease  *string
}

func (p ExpeditedWindowsQualityUpdateSettingsOperationPredicate) Matches(input ExpeditedWindowsQualityUpdateSettings) bool {

	if p.DaysUntilForcedReboot != nil && (input.DaysUntilForcedReboot == nil || *p.DaysUntilForcedReboot != *input.DaysUntilForcedReboot) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QualityUpdateRelease != nil && (input.QualityUpdateRelease == nil || *p.QualityUpdateRelease != *input.QualityUpdateRelease) {
		return false
	}

	return true
}
