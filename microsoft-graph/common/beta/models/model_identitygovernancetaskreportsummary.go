package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportSummary struct {
	FailedTasks      *int64  `json:"failedTasks,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	SuccessfulTasks  *int64  `json:"successfulTasks,omitempty"`
	TotalTasks       *int64  `json:"totalTasks,omitempty"`
	UnprocessedTasks *int64  `json:"unprocessedTasks,omitempty"`
}
