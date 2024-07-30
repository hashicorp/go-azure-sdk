package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleAssignmentRequest struct {
	AssignmentState   *string             `json:"assignmentState,omitempty"`
	Duration          *string             `json:"duration,omitempty"`
	Id                *string             `json:"id,omitempty"`
	ODataType         *string             `json:"@odata.type,omitempty"`
	Reason            *string             `json:"reason,omitempty"`
	RequestedDateTime *string             `json:"requestedDateTime,omitempty"`
	RoleId            *string             `json:"roleId,omitempty"`
	RoleInfo          *PrivilegedRole     `json:"roleInfo,omitempty"`
	Schedule          *GovernanceSchedule `json:"schedule,omitempty"`
	Status            *string             `json:"status,omitempty"`
	TicketNumber      *string             `json:"ticketNumber,omitempty"`
	TicketSystem      *string             `json:"ticketSystem,omitempty"`
	Type              *string             `json:"type,omitempty"`
	UserId            *string             `json:"userId,omitempty"`
}
