package checkprincipalaccess

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckAccessDecision struct {
	AccessDecision *string                `json:"accessDecision,omitempty"`
	ActionId       *string                `json:"actionId,omitempty"`
	RoleAssignment *RoleAssignmentDetails `json:"roleAssignment,omitempty"`
}
