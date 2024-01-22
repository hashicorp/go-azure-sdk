package billingaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceSectionWithCreateSubPermissionOperationPredicate struct {
	BillingProfileDisplayName *string
	BillingProfileId          *string
	InvoiceSectionDisplayName *string
	InvoiceSectionId          *string
}

func (p InvoiceSectionWithCreateSubPermissionOperationPredicate) Matches(input InvoiceSectionWithCreateSubPermission) bool {

	if p.BillingProfileDisplayName != nil && (input.BillingProfileDisplayName == nil || *p.BillingProfileDisplayName != *input.BillingProfileDisplayName) {
		return false
	}

	if p.BillingProfileId != nil && (input.BillingProfileId == nil || *p.BillingProfileId != *input.BillingProfileId) {
		return false
	}

	if p.InvoiceSectionDisplayName != nil && (input.InvoiceSectionDisplayName == nil || *p.InvoiceSectionDisplayName != *input.InvoiceSectionDisplayName) {
		return false
	}

	if p.InvoiceSectionId != nil && (input.InvoiceSectionId == nil || *p.InvoiceSectionId != *input.InvoiceSectionId) {
		return false
	}

	return true
}
