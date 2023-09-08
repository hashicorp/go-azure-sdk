package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelingOptionsOperationPredicate struct {
	LabelId   *string
	ODataType *string
}

func (p SecurityLabelingOptionsOperationPredicate) Matches(input SecurityLabelingOptions) bool {

	if p.LabelId != nil && (input.LabelId == nil || *p.LabelId != *input.LabelId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
