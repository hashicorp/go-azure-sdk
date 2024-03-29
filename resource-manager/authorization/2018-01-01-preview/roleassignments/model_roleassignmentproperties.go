package roleassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentProperties struct {
	CanDelegate      *bool  `json:"canDelegate,omitempty"`
	PrincipalId      string `json:"principalId"`
	RoleDefinitionId string `json:"roleDefinitionId"`
}
