package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReport struct {
	CompletedDateTime     *string                                       `json:"completedDateTime,omitempty"`
	FailedUsersCount      *int64                                        `json:"failedUsersCount,omitempty"`
	Id                    *string                                       `json:"id,omitempty"`
	LastUpdatedDateTime   *string                                       `json:"lastUpdatedDateTime,omitempty"`
	ODataType             *string                                       `json:"@odata.type,omitempty"`
	ProcessingStatus      *IdentityGovernanceTaskReportProcessingStatus `json:"processingStatus,omitempty"`
	RunId                 *string                                       `json:"runId,omitempty"`
	StartedDateTime       *string                                       `json:"startedDateTime,omitempty"`
	SuccessfulUsersCount  *int64                                        `json:"successfulUsersCount,omitempty"`
	Task                  *IdentityGovernanceTask                       `json:"task,omitempty"`
	TaskDefinition        *IdentityGovernanceTaskDefinition             `json:"taskDefinition,omitempty"`
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult     `json:"taskProcessingResults,omitempty"`
	TotalUsersCount       *int64                                        `json:"totalUsersCount,omitempty"`
	UnprocessedUsersCount *int64                                        `json:"unprocessedUsersCount,omitempty"`
}
