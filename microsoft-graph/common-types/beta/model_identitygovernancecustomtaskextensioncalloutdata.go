package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCalloutData struct {
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	Subject              *User                                   `json:"subject,omitempty"`
	Task                 *IdentityGovernanceTask                 `json:"task,omitempty"`
	TaskProcessingresult *IdentityGovernanceTaskProcessingResult `json:"taskProcessingresult,omitempty"`
	Workflow             *IdentityGovernanceWorkflow             `json:"workflow,omitempty"`
}
