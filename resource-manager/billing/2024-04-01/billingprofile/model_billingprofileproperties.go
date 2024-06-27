package billingprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileProperties struct {
	BillTo                   *AddressDetails                 `json:"billTo,omitempty"`
	BillingRelationshipType  *BillingRelationshipType        `json:"billingRelationshipType,omitempty"`
	Currency                 *string                         `json:"currency,omitempty"`
	CurrentPaymentTerm       *PaymentTerm                    `json:"currentPaymentTerm,omitempty"`
	DisplayName              *string                         `json:"displayName,omitempty"`
	EnabledAzurePlans        *[]AzurePlan                    `json:"enabledAzurePlans,omitempty"`
	HasReadAccess            *bool                           `json:"hasReadAccess,omitempty"`
	IndirectRelationshipInfo *IndirectRelationshipInfo       `json:"indirectRelationshipInfo,omitempty"`
	InvoiceDay               *int64                          `json:"invoiceDay,omitempty"`
	InvoiceEmailOptIn        *bool                           `json:"invoiceEmailOptIn,omitempty"`
	InvoiceRecipients        *[]string                       `json:"invoiceRecipients,omitempty"`
	OtherPaymentTerms        *[]PaymentTerm                  `json:"otherPaymentTerms,omitempty"`
	PoNumber                 *string                         `json:"poNumber,omitempty"`
	ProvisioningState        *ProvisioningState              `json:"provisioningState,omitempty"`
	ShipTo                   *AddressDetails                 `json:"shipTo,omitempty"`
	SoldTo                   *AddressDetails                 `json:"soldTo,omitempty"`
	SpendingLimit            *SpendingLimit                  `json:"spendingLimit,omitempty"`
	SpendingLimitDetails     *[]SpendingLimitDetails         `json:"spendingLimitDetails,omitempty"`
	Status                   *BillingProfileStatus           `json:"status,omitempty"`
	StatusReasonCode         *BillingProfileStatusReasonCode `json:"statusReasonCode,omitempty"`
	SystemId                 *string                         `json:"systemId,omitempty"`
	Tags                     *map[string]string              `json:"tags,omitempty"`
	TargetClouds             *[]TargetCloud                  `json:"targetClouds,omitempty"`
}
