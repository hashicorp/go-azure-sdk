package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleAssignment struct {
	AssignmentState                *string                   `json:"assignmentState,omitempty"`
	EndDateTime                    *string                   `json:"endDateTime,omitempty"`
	ExternalId                     *string                   `json:"externalId,omitempty"`
	Id                             *string                   `json:"id,omitempty"`
	LinkedEligibleRoleAssignment   *GovernanceRoleAssignment `json:"linkedEligibleRoleAssignment,omitempty"`
	LinkedEligibleRoleAssignmentId *string                   `json:"linkedEligibleRoleAssignmentId,omitempty"`
	MemberType                     *string                   `json:"memberType,omitempty"`
	ODataType                      *string                   `json:"@odata.type,omitempty"`
	Resource                       *GovernanceResource       `json:"resource,omitempty"`
	ResourceId                     *string                   `json:"resourceId,omitempty"`
	RoleDefinition                 *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`
	RoleDefinitionId               *string                   `json:"roleDefinitionId,omitempty"`
	StartDateTime                  *string                   `json:"startDateTime,omitempty"`
	Status                         *string                   `json:"status,omitempty"`
	Subject                        *GovernanceSubject        `json:"subject,omitempty"`
	SubjectId                      *string                   `json:"subjectId,omitempty"`
}
