package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetails struct {
	DesiredRole  *MultiTenantOrganizationMemberTransitionDetailsDesiredRole  `json:"desiredRole,omitempty"`
	DesiredState *MultiTenantOrganizationMemberTransitionDetailsDesiredState `json:"desiredState,omitempty"`
	Details      *string                                                     `json:"details,omitempty"`
	ODataType    *string                                                     `json:"@odata.type,omitempty"`
	Status       *MultiTenantOrganizationMemberTransitionDetailsStatus       `json:"status,omitempty"`
}
