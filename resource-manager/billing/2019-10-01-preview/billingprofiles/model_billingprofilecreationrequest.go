package billingprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileCreationRequest struct {
	Address           *AddressDetails `json:"address,omitempty"`
	DisplayName       *string         `json:"displayName,omitempty"`
	EnabledAzurePlans *[]AzurePlan    `json:"enabledAzurePlans,omitempty"`
	InvoiceEmailOptIn *bool           `json:"invoiceEmailOptIn,omitempty"`
	PoNumber          *string         `json:"poNumber,omitempty"`
}
