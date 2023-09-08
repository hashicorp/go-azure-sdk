package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomerPayment struct {
	Amount                 *float64  `json:"amount,omitempty"`
	AppliesToInvoiceId     *string   `json:"appliesToInvoiceId,omitempty"`
	AppliesToInvoiceNumber *string   `json:"appliesToInvoiceNumber,omitempty"`
	Comment                *string   `json:"comment,omitempty"`
	ContactId              *string   `json:"contactId,omitempty"`
	Customer               *Customer `json:"customer,omitempty"`
	CustomerId             *string   `json:"customerId,omitempty"`
	CustomerNumber         *string   `json:"customerNumber,omitempty"`
	Description            *string   `json:"description,omitempty"`
	DocumentNumber         *string   `json:"documentNumber,omitempty"`
	ExternalDocumentNumber *string   `json:"externalDocumentNumber,omitempty"`
	Id                     *string   `json:"id,omitempty"`
	JournalDisplayName     *string   `json:"journalDisplayName,omitempty"`
	LastModifiedDateTime   *string   `json:"lastModifiedDateTime,omitempty"`
	LineNumber             *int64    `json:"lineNumber,omitempty"`
	ODataType              *string   `json:"@odata.type,omitempty"`
	PostingDate            *string   `json:"postingDate,omitempty"`
}
