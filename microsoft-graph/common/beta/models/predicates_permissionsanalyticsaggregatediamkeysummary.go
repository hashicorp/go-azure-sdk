package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticsAggregatedIamKeySummaryOperationPredicate struct {
	FindingsCountOverLimit *int64
	ODataType              *string
	TotalCount             *int64
}

func (p PermissionsAnalyticsAggregatedIamKeySummaryOperationPredicate) Matches(input PermissionsAnalyticsAggregatedIamKeySummary) bool {

	if p.FindingsCountOverLimit != nil && (input.FindingsCountOverLimit == nil || *p.FindingsCountOverLimit != *input.FindingsCountOverLimit) {
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
