package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpEvaluationWindowsDevicesInputOperationPredicate struct {
	ODataType *string
	SharedBy  *string
}

func (p DlpEvaluationWindowsDevicesInputOperationPredicate) Matches(input DlpEvaluationWindowsDevicesInput) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SharedBy != nil && (input.SharedBy == nil || *p.SharedBy != *input.SharedBy) {
		return false
	}

	return true
}
