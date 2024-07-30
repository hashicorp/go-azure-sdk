package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTask struct {
	Arguments             *[]KeyValuePair                           `json:"arguments,omitempty"`
	Category              *IdentityGovernanceTaskCategory           `json:"category,omitempty"`
	ContinueOnError       *bool                                     `json:"continueOnError,omitempty"`
	Description           *string                                   `json:"description,omitempty"`
	DisplayName           *string                                   `json:"displayName,omitempty"`
	ExecutionSequence     *int64                                    `json:"executionSequence,omitempty"`
	Id                    *string                                   `json:"id,omitempty"`
	IsEnabled             *bool                                     `json:"isEnabled,omitempty"`
	ODataType             *string                                   `json:"@odata.type,omitempty"`
	TaskDefinitionId      *string                                   `json:"taskDefinitionId,omitempty"`
	TaskProcessingResults *[]IdentityGovernanceTaskProcessingResult `json:"taskProcessingResults,omitempty"`
}
