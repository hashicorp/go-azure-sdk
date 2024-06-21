package transfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransferProperties struct {
	BillingAccountId       *string                   `json:"billingAccountId,omitempty"`
	BillingProfileId       *string                   `json:"billingProfileId,omitempty"`
	CanceledBy             *string                   `json:"canceledBy,omitempty"`
	CreationTime           *string                   `json:"creationTime,omitempty"`
	DetailedTransferStatus *[]DetailedTransferStatus `json:"detailedTransferStatus,omitempty"`
	ExpirationTime         *string                   `json:"expirationTime,omitempty"`
	InitiatorCustomerType  *string                   `json:"initiatorCustomerType,omitempty"`
	InitiatorEmailId       *string                   `json:"initiatorEmailId,omitempty"`
	InvoiceSectionId       *string                   `json:"invoiceSectionId,omitempty"`
	LastModifiedTime       *string                   `json:"lastModifiedTime,omitempty"`
	RecipientEmailId       *string                   `json:"recipientEmailId,omitempty"`
	ResellerId             *string                   `json:"resellerId,omitempty"`
	ResellerName           *string                   `json:"resellerName,omitempty"`
	TransferStatus         *TransferStatus           `json:"transferStatus,omitempty"`
}
