package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTransactionSummaryOperationPredicate struct {
	BlockedCount *int64
	ODataType    *string
	TotalCount   *int64
}

func (p NetworkaccessTransactionSummaryOperationPredicate) Matches(input NetworkaccessTransactionSummary) bool {

	if p.BlockedCount != nil && (input.BlockedCount == nil || *p.BlockedCount != *input.BlockedCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalCount != nil && (input.TotalCount == nil || *p.TotalCount != *input.TotalCount) {
		return false
	}

	return true
}
