package checkpolicyrestrictions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyReference struct {
	PolicyAssignmentId          *string `json:"policyAssignmentId,omitempty"`
	PolicyDefinitionId          *string `json:"policyDefinitionId,omitempty"`
	PolicyDefinitionReferenceId *string `json:"policyDefinitionReferenceId,omitempty"`
	PolicySetDefinitionId       *string `json:"policySetDefinitionId,omitempty"`
}
