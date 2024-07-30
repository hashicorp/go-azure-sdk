package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequest struct {
	AccessId          *PrivilegedAccessGroupAssignmentScheduleRequestAccessId `json:"accessId,omitempty"`
	Action            *PrivilegedAccessGroupAssignmentScheduleRequestAction   `json:"action,omitempty"`
	ActivatedUsing    *PrivilegedAccessGroupEligibilitySchedule               `json:"activatedUsing,omitempty"`
	ApprovalId        *string                                                 `json:"approvalId,omitempty"`
	CompletedDateTime *string                                                 `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet                                            `json:"createdBy,omitempty"`
	CreatedDateTime   *string                                                 `json:"createdDateTime,omitempty"`
	CustomData        *string                                                 `json:"customData,omitempty"`
	Group             *Group                                                  `json:"group,omitempty"`
	GroupId           *string                                                 `json:"groupId,omitempty"`
	Id                *string                                                 `json:"id,omitempty"`
	IsValidationOnly  *bool                                                   `json:"isValidationOnly,omitempty"`
	Justification     *string                                                 `json:"justification,omitempty"`
	ODataType         *string                                                 `json:"@odata.type,omitempty"`
	Principal         *DirectoryObject                                        `json:"principal,omitempty"`
	PrincipalId       *string                                                 `json:"principalId,omitempty"`
	ScheduleInfo      *RequestSchedule                                        `json:"scheduleInfo,omitempty"`
	Status            *string                                                 `json:"status,omitempty"`
	TargetSchedule    *PrivilegedAccessGroupEligibilitySchedule               `json:"targetSchedule,omitempty"`
	TargetScheduleId  *string                                                 `json:"targetScheduleId,omitempty"`
	TicketInfo        *TicketInfo                                             `json:"ticketInfo,omitempty"`
}
