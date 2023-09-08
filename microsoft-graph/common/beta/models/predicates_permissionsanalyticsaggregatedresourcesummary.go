package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticsAggregatedResourceSummaryOperationPredicate struct {
	FindingsCount *int64
	ODataType     *string
	TotalCount    *int64
}

func (p PermissionsAnalyticsAggregatedResourceSummaryOperationPredicate) Matches(input PermissionsAnalyticsAggregatedResourceSummary) bool {

	if p.FindingsCount != nil && (input.FindingsCount == nil || *p.FindingsCount != *input.FindingsCount) {
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
