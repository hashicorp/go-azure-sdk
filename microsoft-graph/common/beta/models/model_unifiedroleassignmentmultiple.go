package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentMultiple struct {
	AppScopeIds       *[]string              `json:"appScopeIds,omitempty"`
	AppScopes         *[]AppScope            `json:"appScopes,omitempty"`
	Condition         *string                `json:"condition,omitempty"`
	Description       *string                `json:"description,omitempty"`
	DirectoryScopeIds *[]string              `json:"directoryScopeIds,omitempty"`
	DirectoryScopes   *[]DirectoryObject     `json:"directoryScopes,omitempty"`
	DisplayName       *string                `json:"displayName,omitempty"`
	Id                *string                `json:"id,omitempty"`
	ODataType         *string                `json:"@odata.type,omitempty"`
	PrincipalIds      *[]string              `json:"principalIds,omitempty"`
	Principals        *[]DirectoryObject     `json:"principals,omitempty"`
	RoleDefinition    *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
	RoleDefinitionId  *string                `json:"roleDefinitionId,omitempty"`
}
