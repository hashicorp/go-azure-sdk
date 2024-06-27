package associatedtenant

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociatedTenantProperties struct {
	BillingManagementState       *BillingManagementTenantState `json:"billingManagementState,omitempty"`
	DisplayName                  *string                       `json:"displayName,omitempty"`
	ProvisioningBillingRequestId *string                       `json:"provisioningBillingRequestId,omitempty"`
	ProvisioningManagementState  *ProvisioningTenantState      `json:"provisioningManagementState,omitempty"`
	ProvisioningState            *ProvisioningState            `json:"provisioningState,omitempty"`
	TenantId                     *string                       `json:"tenantId,omitempty"`
}
