package invoices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvoiceRebillDetails struct {
	LatestInvoiceId    *string             `json:"latestInvoiceId,omitempty"`
	RebillDocumentType *RebillDocumentType `json:"rebillDocumentType,omitempty"`
	RebilledInvoiceId  *string             `json:"rebilledInvoiceId,omitempty"`
}
