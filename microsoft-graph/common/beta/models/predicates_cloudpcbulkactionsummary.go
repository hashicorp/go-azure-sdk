package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkActionSummaryOperationPredicate struct {
	FailedCount       *int64
	InProgressCount   *int64
	NotSupportedCount *int64
	ODataType         *string
	PendingCount      *int64
	SuccessfulCount   *int64
}

func (p CloudPCBulkActionSummaryOperationPredicate) Matches(input CloudPCBulkActionSummary) bool {

	if p.FailedCount != nil && (input.FailedCount == nil || *p.FailedCount != *input.FailedCount) {
		return false
	}

	if p.InProgressCount != nil && (input.InProgressCount == nil || *p.InProgressCount != *input.InProgressCount) {
		return false
	}

	if p.NotSupportedCount != nil && (input.NotSupportedCount == nil || *p.NotSupportedCount != *input.NotSupportedCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PendingCount != nil && (input.PendingCount == nil || *p.PendingCount != *input.PendingCount) {
		return false
	}

	if p.SuccessfulCount != nil && (input.SuccessfulCount == nil || *p.SuccessfulCount != *input.SuccessfulCount) {
		return false
	}

	return true
}
