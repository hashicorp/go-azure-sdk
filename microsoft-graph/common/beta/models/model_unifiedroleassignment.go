package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignment struct {
	AppScope                *AppScope              `json:"appScope,omitempty"`
	AppScopeId              *string                `json:"appScopeId,omitempty"`
	Condition               *string                `json:"condition,omitempty"`
	DirectoryScope          *DirectoryObject       `json:"directoryScope,omitempty"`
	DirectoryScopeId        *string                `json:"directoryScopeId,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	ODataType               *string                `json:"@odata.type,omitempty"`
	Principal               *DirectoryObject       `json:"principal,omitempty"`
	PrincipalId             *string                `json:"principalId,omitempty"`
	PrincipalOrganizationId *string                `json:"principalOrganizationId,omitempty"`
	ResourceScope           *string                `json:"resourceScope,omitempty"`
	RoleDefinition          *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
	RoleDefinitionId        *string                `json:"roleDefinitionId,omitempty"`
}
