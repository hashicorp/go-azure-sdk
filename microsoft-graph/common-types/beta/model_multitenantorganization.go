package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganization struct {
	CreatedDateTime *string                                   `json:"createdDateTime,omitempty"`
	Description     *string                                   `json:"description,omitempty"`
	DisplayName     *string                                   `json:"displayName,omitempty"`
	Id              *string                                   `json:"id,omitempty"`
	JoinRequest     *MultiTenantOrganizationJoinRequestRecord `json:"joinRequest,omitempty"`
	ODataType       *string                                   `json:"@odata.type,omitempty"`
	State           *MultiTenantOrganizationState             `json:"state,omitempty"`
	Tenants         *[]MultiTenantOrganizationMember          `json:"tenants,omitempty"`
}
