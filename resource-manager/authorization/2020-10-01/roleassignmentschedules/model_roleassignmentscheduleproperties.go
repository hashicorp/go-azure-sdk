package roleassignmentschedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleProperties struct {
	AssignmentType                  *AssignmentType     `json:"assignmentType,omitempty"`
	Condition                       *string             `json:"condition,omitempty"`
	ConditionVersion                *string             `json:"conditionVersion,omitempty"`
	CreatedOn                       *string             `json:"createdOn,omitempty"`
	EndDateTime                     *string             `json:"endDateTime,omitempty"`
	ExpandedProperties              *ExpandedProperties `json:"expandedProperties,omitempty"`
	LinkedRoleEligibilityScheduleId *string             `json:"linkedRoleEligibilityScheduleId,omitempty"`
	MemberType                      *MemberType         `json:"memberType,omitempty"`
	PrincipalId                     *string             `json:"principalId,omitempty"`
	PrincipalType                   *PrincipalType      `json:"principalType,omitempty"`
	RoleAssignmentScheduleRequestId *string             `json:"roleAssignmentScheduleRequestId,omitempty"`
	RoleDefinitionId                *string             `json:"roleDefinitionId,omitempty"`
	Scope                           *string             `json:"scope,omitempty"`
	StartDateTime                   *string             `json:"startDateTime,omitempty"`
	Status                          *Status             `json:"status,omitempty"`
	UpdatedOn                       *string             `json:"updatedOn,omitempty"`
}
