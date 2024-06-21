package denyassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DenyAssignmentProperties struct {
	Condition               *string                     `json:"condition,omitempty"`
	ConditionVersion        *string                     `json:"conditionVersion,omitempty"`
	CreatedBy               *string                     `json:"createdBy,omitempty"`
	CreatedOn               *string                     `json:"createdOn,omitempty"`
	DenyAssignmentName      *string                     `json:"denyAssignmentName,omitempty"`
	Description             *string                     `json:"description,omitempty"`
	DoNotApplyToChildScopes *bool                       `json:"doNotApplyToChildScopes,omitempty"`
	ExcludePrincipals       *[]Principal                `json:"excludePrincipals,omitempty"`
	IsSystemProtected       *bool                       `json:"isSystemProtected,omitempty"`
	Permissions             *[]DenyAssignmentPermission `json:"permissions,omitempty"`
	Principals              *[]Principal                `json:"principals,omitempty"`
	Scope                   *string                     `json:"scope,omitempty"`
	UpdatedBy               *string                     `json:"updatedBy,omitempty"`
	UpdatedOn               *string                     `json:"updatedOn,omitempty"`
}
