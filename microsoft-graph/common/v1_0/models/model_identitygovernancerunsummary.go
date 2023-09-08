package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunSummary struct {
	FailedRuns     *int64  `json:"failedRuns,omitempty"`
	FailedTasks    *int64  `json:"failedTasks,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	SuccessfulRuns *int64  `json:"successfulRuns,omitempty"`
	TotalRuns      *int64  `json:"totalRuns,omitempty"`
	TotalTasks     *int64  `json:"totalTasks,omitempty"`
	TotalUsers     *int64  `json:"totalUsers,omitempty"`
}
