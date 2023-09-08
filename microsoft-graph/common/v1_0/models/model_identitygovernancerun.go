package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRun struct {
	CompletedDateTime          *string                                     `json:"completedDateTime,omitempty"`
	FailedTasksCount           *int64                                      `json:"failedTasksCount,omitempty"`
	FailedUsersCount           *int64                                      `json:"failedUsersCount,omitempty"`
	Id                         *string                                     `json:"id,omitempty"`
	LastUpdatedDateTime        *string                                     `json:"lastUpdatedDateTime,omitempty"`
	ODataType                  *string                                     `json:"@odata.type,omitempty"`
	ProcessingStatus           *IdentityGovernanceRunProcessingStatus      `json:"processingStatus,omitempty"`
	ScheduledDateTime          *string                                     `json:"scheduledDateTime,omitempty"`
	StartedDateTime            *string                                     `json:"startedDateTime,omitempty"`
	SuccessfulUsersCount       *int64                                      `json:"successfulUsersCount,omitempty"`
	TaskProcessingResults      *[]IdentityGovernanceTaskProcessingResult   `json:"taskProcessingResults,omitempty"`
	TotalTasksCount            *int64                                      `json:"totalTasksCount,omitempty"`
	TotalUnprocessedTasksCount *int64                                      `json:"totalUnprocessedTasksCount,omitempty"`
	TotalUsersCount            *int64                                      `json:"totalUsersCount,omitempty"`
	UserProcessingResults      *[]IdentityGovernanceUserProcessingResult   `json:"userProcessingResults,omitempty"`
	WorkflowExecutionType      *IdentityGovernanceRunWorkflowExecutionType `json:"workflowExecutionType,omitempty"`
}
