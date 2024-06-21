package roleassignmentschedulerequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleRequestProperties struct {
	ApprovalId                             *string                                              `json:"approvalId,omitempty"`
	Condition                              *string                                              `json:"condition,omitempty"`
	ConditionVersion                       *string                                              `json:"conditionVersion,omitempty"`
	CreatedOn                              *string                                              `json:"createdOn,omitempty"`
	ExpandedProperties                     *ExpandedProperties                                  `json:"expandedProperties,omitempty"`
	Justification                          *string                                              `json:"justification,omitempty"`
	LinkedRoleEligibilityScheduleId        *string                                              `json:"linkedRoleEligibilityScheduleId,omitempty"`
	PrincipalId                            string                                               `json:"principalId"`
	PrincipalType                          *PrincipalType                                       `json:"principalType,omitempty"`
	RequestType                            RequestType                                          `json:"requestType"`
	RequestorId                            *string                                              `json:"requestorId,omitempty"`
	RoleDefinitionId                       string                                               `json:"roleDefinitionId"`
	ScheduleInfo                           *RoleAssignmentScheduleRequestPropertiesScheduleInfo `json:"scheduleInfo,omitempty"`
	Scope                                  *string                                              `json:"scope,omitempty"`
	Status                                 *Status                                              `json:"status,omitempty"`
	TargetRoleAssignmentScheduleId         *string                                              `json:"targetRoleAssignmentScheduleId,omitempty"`
	TargetRoleAssignmentScheduleInstanceId *string                                              `json:"targetRoleAssignmentScheduleInstanceId,omitempty"`
	TicketInfo                             *RoleAssignmentScheduleRequestPropertiesTicketInfo   `json:"ticketInfo,omitempty"`
}
