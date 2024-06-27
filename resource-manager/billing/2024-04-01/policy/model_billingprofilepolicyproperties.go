package policy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfilePolicyProperties struct {
	EnterpriseAgreementPolicies   *EnterpriseAgreementPolicies         `json:"enterpriseAgreementPolicies,omitempty"`
	InvoiceSectionLabelManagement *InvoiceSectionLabelManagementPolicy `json:"invoiceSectionLabelManagement,omitempty"`
	MarketplacePurchases          *MarketplacePurchasesPolicy          `json:"marketplacePurchases,omitempty"`
	Policies                      *[]PolicySummary                     `json:"policies,omitempty"`
	ProvisioningState             *ProvisioningState                   `json:"provisioningState,omitempty"`
	ReservationPurchases          *ReservationPurchasesPolicy          `json:"reservationPurchases,omitempty"`
	SavingsPlanPurchases          *SavingsPlanPurchasesPolicy          `json:"savingsPlanPurchases,omitempty"`
	ViewCharges                   *ViewChargesPolicy                   `json:"viewCharges,omitempty"`
}
