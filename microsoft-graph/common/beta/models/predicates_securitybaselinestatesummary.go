package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineStateSummaryOperationPredicate struct {
	ConflictCount      *int64
	ErrorCount         *int64
	Id                 *string
	NotApplicableCount *int64
	NotSecureCount     *int64
	ODataType          *string
	SecureCount        *int64
	UnknownCount       *int64
}

func (p SecurityBaselineStateSummaryOperationPredicate) Matches(input SecurityBaselineStateSummary) bool {

	if p.ConflictCount != nil && (input.ConflictCount == nil || *p.ConflictCount != *input.ConflictCount) {
		return false
	}

	if p.ErrorCount != nil && (input.ErrorCount == nil || *p.ErrorCount != *input.ErrorCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NotApplicableCount != nil && (input.NotApplicableCount == nil || *p.NotApplicableCount != *input.NotApplicableCount) {
		return false
	}

	if p.NotSecureCount != nil && (input.NotSecureCount == nil || *p.NotSecureCount != *input.NotSecureCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SecureCount != nil && (input.SecureCount == nil || *p.SecureCount != *input.SecureCount) {
		return false
	}

	if p.UnknownCount != nil && (input.UnknownCount == nil || *p.UnknownCount != *input.UnknownCount) {
		return false
	}

	return true
}
