package roleassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentProperties struct {
	Condition                          *string        `json:"condition,omitempty"`
	ConditionVersion                   *string        `json:"conditionVersion,omitempty"`
	CreatedBy                          *string        `json:"createdBy,omitempty"`
	CreatedOn                          *string        `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string        `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string        `json:"description,omitempty"`
	PrincipalId                        string         `json:"principalId"`
	PrincipalType                      *PrincipalType `json:"principalType,omitempty"`
	RoleDefinitionId                   string         `json:"roleDefinitionId"`
	Scope                              *string        `json:"scope,omitempty"`
	UpdatedBy                          *string        `json:"updatedBy,omitempty"`
	UpdatedOn                          *string        `json:"updatedOn,omitempty"`
}
