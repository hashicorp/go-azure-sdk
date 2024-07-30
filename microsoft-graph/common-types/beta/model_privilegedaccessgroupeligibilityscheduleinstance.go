package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstance struct {
	AccessId              *PrivilegedAccessGroupEligibilityScheduleInstanceAccessId   `json:"accessId,omitempty"`
	EligibilityScheduleId *string                                                     `json:"eligibilityScheduleId,omitempty"`
	EndDateTime           *string                                                     `json:"endDateTime,omitempty"`
	Group                 *Group                                                      `json:"group,omitempty"`
	GroupId               *string                                                     `json:"groupId,omitempty"`
	Id                    *string                                                     `json:"id,omitempty"`
	MemberType            *PrivilegedAccessGroupEligibilityScheduleInstanceMemberType `json:"memberType,omitempty"`
	ODataType             *string                                                     `json:"@odata.type,omitempty"`
	Principal             *DirectoryObject                                            `json:"principal,omitempty"`
	PrincipalId           *string                                                     `json:"principalId,omitempty"`
	StartDateTime         *string                                                     `json:"startDateTime,omitempty"`
}
