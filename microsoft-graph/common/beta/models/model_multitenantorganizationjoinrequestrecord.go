package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecord struct {
	AddedByTenantId   *string                                              `json:"addedByTenantId,omitempty"`
	Id                *string                                              `json:"id,omitempty"`
	MemberState       *MultiTenantOrganizationJoinRequestRecordMemberState `json:"memberState,omitempty"`
	ODataType         *string                                              `json:"@odata.type,omitempty"`
	Role              *MultiTenantOrganizationJoinRequestRecordRole        `json:"role,omitempty"`
	TransitionDetails *MultiTenantOrganizationJoinRequestTransitionDetails `json:"transitionDetails,omitempty"`
}
