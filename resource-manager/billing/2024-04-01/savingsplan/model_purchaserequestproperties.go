package savingsplan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PurchaseRequestProperties struct {
	AppliedScopeProperties *AppliedScopeProperties `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType       *AppliedScopeType       `json:"appliedScopeType,omitempty"`
	BillingPlan            *BillingPlan            `json:"billingPlan,omitempty"`
	BillingScopeId         *string                 `json:"billingScopeId,omitempty"`
	Commitment             *Commitment             `json:"commitment,omitempty"`
	DisplayName            *string                 `json:"displayName,omitempty"`
	Renew                  *bool                   `json:"renew,omitempty"`
	Term                   *SavingsPlanTerm        `json:"term,omitempty"`
}
