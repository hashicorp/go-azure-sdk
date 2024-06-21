package roleeligibilityschedulerequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleRequestProperties struct {
	ApprovalId                              *string                                               `json:"approvalId,omitempty"`
	Condition                               *string                                               `json:"condition,omitempty"`
	ConditionVersion                        *string                                               `json:"conditionVersion,omitempty"`
	CreatedOn                               *string                                               `json:"createdOn,omitempty"`
	ExpandedProperties                      *ExpandedProperties                                   `json:"expandedProperties,omitempty"`
	Justification                           *string                                               `json:"justification,omitempty"`
	PrincipalId                             string                                                `json:"principalId"`
	PrincipalType                           *PrincipalType                                        `json:"principalType,omitempty"`
	RequestType                             RequestType                                           `json:"requestType"`
	RequestorId                             *string                                               `json:"requestorId,omitempty"`
	RoleDefinitionId                        string                                                `json:"roleDefinitionId"`
	ScheduleInfo                            *RoleEligibilityScheduleRequestPropertiesScheduleInfo `json:"scheduleInfo,omitempty"`
	Scope                                   *string                                               `json:"scope,omitempty"`
	Status                                  *Status                                               `json:"status,omitempty"`
	TargetRoleEligibilityScheduleId         *string                                               `json:"targetRoleEligibilityScheduleId,omitempty"`
	TargetRoleEligibilityScheduleInstanceId *string                                               `json:"targetRoleEligibilityScheduleInstanceId,omitempty"`
	TicketInfo                              *RoleEligibilityScheduleRequestPropertiesTicketInfo   `json:"ticketInfo,omitempty"`
}
