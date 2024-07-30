package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleAssignmentRequest struct {
	AssignmentState                *string                                `json:"assignmentState,omitempty"`
	Id                             *string                                `json:"id,omitempty"`
	LinkedEligibleRoleAssignmentId *string                                `json:"linkedEligibleRoleAssignmentId,omitempty"`
	ODataType                      *string                                `json:"@odata.type,omitempty"`
	Reason                         *string                                `json:"reason,omitempty"`
	RequestedDateTime              *string                                `json:"requestedDateTime,omitempty"`
	Resource                       *GovernanceResource                    `json:"resource,omitempty"`
	ResourceId                     *string                                `json:"resourceId,omitempty"`
	RoleDefinition                 *GovernanceRoleDefinition              `json:"roleDefinition,omitempty"`
	RoleDefinitionId               *string                                `json:"roleDefinitionId,omitempty"`
	Schedule                       *GovernanceSchedule                    `json:"schedule,omitempty"`
	Status                         *GovernanceRoleAssignmentRequestStatus `json:"status,omitempty"`
	Subject                        *GovernanceSubject                     `json:"subject,omitempty"`
	SubjectId                      *string                                `json:"subjectId,omitempty"`
	Type                           *string                                `json:"type,omitempty"`
}
