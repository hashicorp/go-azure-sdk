package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseProcessingStateOperationPredicate struct {
	ODataType *string
	State     *string
}

func (p LicenseProcessingStateOperationPredicate) Matches(input LicenseProcessingState) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	return true
}
