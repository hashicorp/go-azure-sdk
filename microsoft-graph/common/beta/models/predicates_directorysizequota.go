package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectorySizeQuotaOperationPredicate struct {
	ODataType *string
	Total     *int64
	Used      *int64
}

func (p DirectorySizeQuotaOperationPredicate) Matches(input DirectorySizeQuota) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
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
