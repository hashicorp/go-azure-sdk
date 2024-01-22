package billingprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileProperties struct {
	Address                  *AddressDetails           `json:"address,omitempty"`
	BillingRelationshipType  *BillingRelationshipType  `json:"billingRelationshipType,omitempty"`
	Currency                 *string                   `json:"currency,omitempty"`
	DisplayName              *string                   `json:"displayName,omitempty"`
	EnabledAzurePlans        *[]AzurePlan              `json:"enabledAzurePlans,omitempty"`
	IndirectRelationshipInfo *IndirectRelationshipInfo `json:"indirectRelationshipInfo,omitempty"`
	InvoiceDay               *int64                    `json:"invoiceDay,omitempty"`
	InvoiceEmailOptIn        *bool                     `json:"invoiceEmailOptIn,omitempty"`
	InvoiceSections          *[]InvoiceSection         `json:"invoiceSections,omitempty"`
	PoNumber                 *string                   `json:"poNumber,omitempty"`
	SpendingLimit            *SpendingLimit            `json:"spendingLimit,omitempty"`
	Status                   *BillingProfileStatus     `json:"status,omitempty"`
	StatusReasonCode         *StatusReasonCode         `json:"statusReasonCode,omitempty"`
	TargetClouds             *[]TargetCloud            `json:"targetClouds,omitempty"`
}
