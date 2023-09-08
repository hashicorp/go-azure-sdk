package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticsAggregatedIamKeySummary struct {
	FindingsCountOverLimit *int64  `json:"findingsCountOverLimit,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	TotalCount             *int64  `json:"totalCount,omitempty"`
}
