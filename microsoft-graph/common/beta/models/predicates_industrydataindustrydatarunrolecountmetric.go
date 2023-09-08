package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunRoleCountMetricOperationPredicate struct {
	Count     *int64
	ODataType *string
	Role      *string
}

func (p IndustryDataIndustryDataRunRoleCountMetricOperationPredicate) Matches(input IndustryDataIndustryDataRunRoleCountMetric) bool {

	if p.Count != nil && (input.Count == nil || *p.Count != *input.Count) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Role != nil && (input.Role == nil || *p.Role != *input.Role) {
		return false
	}

	return true
}
