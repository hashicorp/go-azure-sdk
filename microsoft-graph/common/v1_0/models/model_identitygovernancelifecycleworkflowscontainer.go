package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceLifecycleWorkflowsContainer struct {
	CustomTaskExtensions *[]IdentityGovernanceCustomTaskExtension       `json:"customTaskExtensions,omitempty"`
	DeletedItems         *DeletedItemContainer                          `json:"deletedItems,omitempty"`
	Id                   *string                                        `json:"id,omitempty"`
	ODataType            *string                                        `json:"@odata.type,omitempty"`
	Settings             *IdentityGovernanceLifecycleManagementSettings `json:"settings,omitempty"`
	TaskDefinitions      *[]IdentityGovernanceTaskDefinition            `json:"taskDefinitions,omitempty"`
	WorkflowTemplates    *[]IdentityGovernanceWorkflowTemplate          `json:"workflowTemplates,omitempty"`
	Workflows            *[]IdentityGovernanceWorkflow                  `json:"workflows,omitempty"`
}
