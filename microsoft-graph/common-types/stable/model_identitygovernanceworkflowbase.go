package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowBase struct {
	Category             *IdentityGovernanceWorkflowBaseCategory        `json:"category,omitempty"`
	CreatedBy            *User                                          `json:"createdBy,omitempty"`
	CreatedDateTime      *string                                        `json:"createdDateTime,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	DisplayName          *string                                        `json:"displayName,omitempty"`
	ExecutionConditions  *IdentityGovernanceWorkflowExecutionConditions `json:"executionConditions,omitempty"`
	IsEnabled            *bool                                          `json:"isEnabled,omitempty"`
	IsSchedulingEnabled  *bool                                          `json:"isSchedulingEnabled,omitempty"`
	LastModifiedBy       *User                                          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                        `json:"@odata.type,omitempty"`
	Tasks                *[]IdentityGovernanceTask                      `json:"tasks,omitempty"`
}
