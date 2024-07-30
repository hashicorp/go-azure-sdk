package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskProcessingResult struct {
	CompletedDateTime *string                                                 `json:"completedDateTime,omitempty"`
	CreatedDateTime   *string                                                 `json:"createdDateTime,omitempty"`
	FailureReason     *string                                                 `json:"failureReason,omitempty"`
	Id                *string                                                 `json:"id,omitempty"`
	ODataType         *string                                                 `json:"@odata.type,omitempty"`
	ProcessingStatus  *IdentityGovernanceTaskProcessingResultProcessingStatus `json:"processingStatus,omitempty"`
	StartedDateTime   *string                                                 `json:"startedDateTime,omitempty"`
	Subject           *User                                                   `json:"subject,omitempty"`
	Task              *IdentityGovernanceTask                                 `json:"task,omitempty"`
}
