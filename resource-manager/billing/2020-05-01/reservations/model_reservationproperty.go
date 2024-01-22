package reservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationProperty struct {
	AppliedScopeType             *string                         `json:"appliedScopeType,omitempty"`
	AppliedScopes                *[]string                       `json:"appliedScopes,omitempty"`
	DisplayName                  *string                         `json:"displayName,omitempty"`
	DisplayProvisioningState     *string                         `json:"displayProvisioningState,omitempty"`
	EffectiveDateTime            *string                         `json:"effectiveDateTime,omitempty"`
	ExpiryDate                   *string                         `json:"expiryDate,omitempty"`
	ProvisioningState            *string                         `json:"provisioningState,omitempty"`
	ProvisioningSubState         *string                         `json:"provisioningSubState,omitempty"`
	Quantity                     *float64                        `json:"quantity,omitempty"`
	Renew                        *bool                           `json:"renew,omitempty"`
	RenewSource                  *string                         `json:"renewSource,omitempty"`
	ReservedResourceType         *string                         `json:"reservedResourceType,omitempty"`
	SkuDescription               *string                         `json:"skuDescription,omitempty"`
	Term                         *string                         `json:"term,omitempty"`
	UserFriendlyAppliedScopeType *string                         `json:"userFriendlyAppliedScopeType,omitempty"`
	UserFriendlyRenewState       *string                         `json:"userFriendlyRenewState,omitempty"`
	Utilization                  *ReservationPropertyUtilization `json:"utilization,omitempty"`
}
