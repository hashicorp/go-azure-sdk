package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetails struct {
	DesiredMemberState *MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState `json:"desiredMemberState,omitempty"`
	Details            *string                                                                `json:"details,omitempty"`
	ODataType          *string                                                                `json:"@odata.type,omitempty"`
	Status             *MultiTenantOrganizationJoinRequestTransitionDetailsStatus             `json:"status,omitempty"`
}
