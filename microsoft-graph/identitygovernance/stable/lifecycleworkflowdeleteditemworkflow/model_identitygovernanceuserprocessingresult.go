package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserProcessingResult struct {
	CompletedDateTime          *string                                                      `json:"completedDateTime,omitempty"`
	FailedTasksCount           *int64                                                       `json:"failedTasksCount,omitempty"`
	Id                         *string                                                      `json:"id,omitempty"`
	ODataType                  *string                                                      `json:"@odata.type,omitempty"`
	ProcessingStatus           *IdentityGovernanceUserProcessingResultProcessingStatus      `json:"processingStatus,omitempty"`
	ScheduledDateTime          *string                                                      `json:"scheduledDateTime,omitempty"`
	StartedDateTime            *string                                                      `json:"startedDateTime,omitempty"`
	Subject                    *User                                                        `json:"subject,omitempty"`
	TaskProcessingResults      *[]IdentityGovernanceTaskProcessingResult                    `json:"taskProcessingResults,omitempty"`
	TotalTasksCount            *int64                                                       `json:"totalTasksCount,omitempty"`
	TotalUnprocessedTasksCount *int64                                                       `json:"totalUnprocessedTasksCount,omitempty"`
	WorkflowExecutionType      *IdentityGovernanceUserProcessingResultWorkflowExecutionType `json:"workflowExecutionType,omitempty"`
	WorkflowVersion            *int64                                                       `json:"workflowVersion,omitempty"`
}
