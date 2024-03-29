package billingaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentPolicies struct {
	AccountOwnerViewCharges    *bool `json:"accountOwnerViewCharges,omitempty"`
	DepartmentAdminViewCharges *bool `json:"departmentAdminViewCharges,omitempty"`
	MarketplacesEnabled        *bool `json:"marketplacesEnabled,omitempty"`
	ReservedInstancesEnabled   *bool `json:"reservedInstancesEnabled,omitempty"`
}
