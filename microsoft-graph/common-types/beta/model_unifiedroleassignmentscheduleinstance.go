package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentScheduleInstance struct {
	ActivatedUsing           *UnifiedRoleEligibilityScheduleInstance `json:"activatedUsing,omitempty"`
	AppScope                 *AppScope                               `json:"appScope,omitempty"`
	AppScopeId               *string                                 `json:"appScopeId,omitempty"`
	AssignmentType           *string                                 `json:"assignmentType,omitempty"`
	DirectoryScope           *DirectoryObject                        `json:"directoryScope,omitempty"`
	DirectoryScopeId         *string                                 `json:"directoryScopeId,omitempty"`
	EndDateTime              *string                                 `json:"endDateTime,omitempty"`
	Id                       *string                                 `json:"id,omitempty"`
	MemberType               *string                                 `json:"memberType,omitempty"`
	ODataType                *string                                 `json:"@odata.type,omitempty"`
	Principal                *DirectoryObject                        `json:"principal,omitempty"`
	PrincipalId              *string                                 `json:"principalId,omitempty"`
	RoleAssignmentOriginId   *string                                 `json:"roleAssignmentOriginId,omitempty"`
	RoleAssignmentScheduleId *string                                 `json:"roleAssignmentScheduleId,omitempty"`
	RoleDefinition           *UnifiedRoleDefinition                  `json:"roleDefinition,omitempty"`
	RoleDefinitionId         *string                                 `json:"roleDefinitionId,omitempty"`
	StartDateTime            *string                                 `json:"startDateTime,omitempty"`
}
