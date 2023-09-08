package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaOperationPredicate struct {
	Deleted   *int64
	ODataType *string
	Remaining *int64
	State     *string
	Total     *int64
	Used      *int64
}

func (p QuotaOperationPredicate) Matches(input Quota) bool {

	if p.Deleted != nil && (input.Deleted == nil || *p.Deleted != *input.Deleted) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Remaining != nil && (input.Remaining == nil || *p.Remaining != *input.Remaining) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	if p.Total != nil && (input.Total == nil || *p.Total != *input.Total) {
		return false
	}

	if p.Used != nil && (input.Used == nil || *p.Used != *input.Used) {
		return false
	}

	return true
}
