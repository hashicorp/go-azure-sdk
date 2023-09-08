package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminReportSettingsOperationPredicate struct {
	DisplayConcealedNames *bool
	Id                    *string
	ODataType             *string
}

func (p AdminReportSettingsOperationPredicate) Matches(input AdminReportSettings) bool {

	if p.DisplayConcealedNames != nil && (input.DisplayConcealedNames == nil || *p.DisplayConcealedNames != *input.DisplayConcealedNames) {
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
