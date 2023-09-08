package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetentionSettingOperationPredicate struct {
	Interval  *string
	ODataType *string
	Period    *string
}

func (p RetentionSettingOperationPredicate) Matches(input RetentionSetting) bool {

	if p.Interval != nil && (input.Interval == nil || *p.Interval != *input.Interval) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Period != nil && (input.Period == nil || *p.Period != *input.Period) {
		return false
	}

	return true
}
