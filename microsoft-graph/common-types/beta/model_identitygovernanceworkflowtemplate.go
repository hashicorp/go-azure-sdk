package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowTemplate struct {
	Category            *IdentityGovernanceWorkflowTemplateCategory    `json:"category,omitempty"`
	Description         *string                                        `json:"description,omitempty"`
	DisplayName         *string                                        `json:"displayName,omitempty"`
	ExecutionConditions *IdentityGovernanceWorkflowExecutionConditions `json:"executionConditions,omitempty"`
	Id                  *string                                        `json:"id,omitempty"`
	ODataType           *string                                        `json:"@odata.type,omitempty"`
	Tasks               *[]IdentityGovernanceTask                      `json:"tasks,omitempty"`
}
