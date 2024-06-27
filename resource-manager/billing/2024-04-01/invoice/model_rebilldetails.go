package invoice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RebillDetails struct {
	CreditNoteDocumentId *string        `json:"creditNoteDocumentId,omitempty"`
	InvoiceDocumentId    *string        `json:"invoiceDocumentId,omitempty"`
	RebillDetails        *RebillDetails `json:"rebillDetails,omitempty"`
}
