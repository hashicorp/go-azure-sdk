package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionDurationInDaysOperationPredicate struct {
	Days      *int64
	ODataType *string
}

func (p SecurityRetentionDurationInDaysOperationPredicate) Matches(input SecurityRetentionDurationInDays) bool {

	if p.Days != nil && (input.Days == nil || *p.Days != *input.Days) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
