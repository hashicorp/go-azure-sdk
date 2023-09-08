package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentScheduleRequest struct {
	Action            *UnifiedRoleAssignmentScheduleRequestAction `json:"action,omitempty"`
	ActivatedUsing    *UnifiedRoleEligibilitySchedule             `json:"activatedUsing,omitempty"`
	AppScope          *AppScope                                   `json:"appScope,omitempty"`
	AppScopeId        *string                                     `json:"appScopeId,omitempty"`
	ApprovalId        *string                                     `json:"approvalId,omitempty"`
	CompletedDateTime *string                                     `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet                                `json:"createdBy,omitempty"`
	CreatedDateTime   *string                                     `json:"createdDateTime,omitempty"`
	CustomData        *string                                     `json:"customData,omitempty"`
	DirectoryScope    *DirectoryObject                            `json:"directoryScope,omitempty"`
	DirectoryScopeId  *string                                     `json:"directoryScopeId,omitempty"`
	Id                *string                                     `json:"id,omitempty"`
	IsValidationOnly  *bool                                       `json:"isValidationOnly,omitempty"`
	Justification     *string                                     `json:"justification,omitempty"`
	ODataType         *string                                     `json:"@odata.type,omitempty"`
	Principal         *DirectoryObject                            `json:"principal,omitempty"`
	PrincipalId       *string                                     `json:"principalId,omitempty"`
	RoleDefinition    *UnifiedRoleDefinition                      `json:"roleDefinition,omitempty"`
	RoleDefinitionId  *string                                     `json:"roleDefinitionId,omitempty"`
	ScheduleInfo      *RequestSchedule                            `json:"scheduleInfo,omitempty"`
	Status            *string                                     `json:"status,omitempty"`
	TargetSchedule    *UnifiedRoleAssignmentSchedule              `json:"targetSchedule,omitempty"`
	TargetScheduleId  *string                                     `json:"targetScheduleId,omitempty"`
	TicketInfo        *TicketInfo                                 `json:"ticketInfo,omitempty"`
}
