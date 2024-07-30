package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentSchedule struct {
	AccessId         *PrivilegedAccessGroupAssignmentScheduleAccessId       `json:"accessId,omitempty"`
	ActivatedUsing   *PrivilegedAccessGroupEligibilitySchedule              `json:"activatedUsing,omitempty"`
	AssignmentType   *PrivilegedAccessGroupAssignmentScheduleAssignmentType `json:"assignmentType,omitempty"`
	CreatedDateTime  *string                                                `json:"createdDateTime,omitempty"`
	CreatedUsing     *string                                                `json:"createdUsing,omitempty"`
	Group            *Group                                                 `json:"group,omitempty"`
	GroupId          *string                                                `json:"groupId,omitempty"`
	Id               *string                                                `json:"id,omitempty"`
	MemberType       *PrivilegedAccessGroupAssignmentScheduleMemberType     `json:"memberType,omitempty"`
	ModifiedDateTime *string                                                `json:"modifiedDateTime,omitempty"`
	ODataType        *string                                                `json:"@odata.type,omitempty"`
	Principal        *DirectoryObject                                       `json:"principal,omitempty"`
	PrincipalId      *string                                                `json:"principalId,omitempty"`
	ScheduleInfo     *RequestSchedule                                       `json:"scheduleInfo,omitempty"`
	Status           *string                                                `json:"status,omitempty"`
}
