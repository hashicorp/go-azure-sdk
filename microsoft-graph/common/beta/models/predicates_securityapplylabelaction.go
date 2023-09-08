package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityApplyLabelActionOperationPredicate struct {
	ODataType          *string
	SensitivityLabelId *string
}

func (p SecurityApplyLabelActionOperationPredicate) Matches(input SecurityApplyLabelAction) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SensitivityLabelId != nil && (input.SensitivityLabelId == nil || *p.SensitivityLabelId != *input.SensitivityLabelId) {
		return false
	}

	return true
}
