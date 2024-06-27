package reservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationAppliedScopeProperties struct {
	DisplayName       *string `json:"displayName,omitempty"`
	ManagementGroupId *string `json:"managementGroupId,omitempty"`
	ResourceGroupId   *string `json:"resourceGroupId,omitempty"`
	SubscriptionId    *string `json:"subscriptionId,omitempty"`
	TenantId          *string `json:"tenantId,omitempty"`
}
