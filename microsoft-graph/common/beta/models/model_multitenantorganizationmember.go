package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMember struct {
	AddedByTenantId   *string                                         `json:"addedByTenantId,omitempty"`
	AddedDateTime     *string                                         `json:"addedDateTime,omitempty"`
	DeletedDateTime   *string                                         `json:"deletedDateTime,omitempty"`
	DisplayName       *string                                         `json:"displayName,omitempty"`
	Id                *string                                         `json:"id,omitempty"`
	JoinedDateTime    *string                                         `json:"joinedDateTime,omitempty"`
	ODataType         *string                                         `json:"@odata.type,omitempty"`
	Role              *MultiTenantOrganizationMemberRole              `json:"role,omitempty"`
	State             *MultiTenantOrganizationMemberState             `json:"state,omitempty"`
	TenantId          *string                                         `json:"tenantId,omitempty"`
	TransitionDetails *MultiTenantOrganizationMemberTransitionDetails `json:"transitionDetails,omitempty"`
}
