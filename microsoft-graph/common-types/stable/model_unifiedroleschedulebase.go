package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleScheduleBase struct {
	AppScope         *AppScope              `json:"appScope,omitempty"`
	AppScopeId       *string                `json:"appScopeId,omitempty"`
	CreatedDateTime  *string                `json:"createdDateTime,omitempty"`
	CreatedUsing     *string                `json:"createdUsing,omitempty"`
	DirectoryScope   *DirectoryObject       `json:"directoryScope,omitempty"`
	DirectoryScopeId *string                `json:"directoryScopeId,omitempty"`
	Id               *string                `json:"id,omitempty"`
	ModifiedDateTime *string                `json:"modifiedDateTime,omitempty"`
	ODataType        *string                `json:"@odata.type,omitempty"`
	Principal        *DirectoryObject       `json:"principal,omitempty"`
	PrincipalId      *string                `json:"principalId,omitempty"`
	RoleDefinition   *UnifiedRoleDefinition `json:"roleDefinition,omitempty"`
	RoleDefinitionId *string                `json:"roleDefinitionId,omitempty"`
	Status           *string                `json:"status,omitempty"`
}
