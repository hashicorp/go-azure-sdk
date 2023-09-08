package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JournalLine struct {
	Account                *Account `json:"account,omitempty"`
	AccountId              *string  `json:"accountId,omitempty"`
	AccountNumber          *string  `json:"accountNumber,omitempty"`
	Amount                 *float64 `json:"amount,omitempty"`
	Comment                *string  `json:"comment,omitempty"`
	Description            *string  `json:"description,omitempty"`
	DocumentNumber         *string  `json:"documentNumber,omitempty"`
	ExternalDocumentNumber *string  `json:"externalDocumentNumber,omitempty"`
	Id                     *string  `json:"id,omitempty"`
	JournalDisplayName     *string  `json:"journalDisplayName,omitempty"`
	LastModifiedDateTime   *string  `json:"lastModifiedDateTime,omitempty"`
	LineNumber             *int64   `json:"lineNumber,omitempty"`
	ODataType              *string  `json:"@odata.type,omitempty"`
	PostingDate            *string  `json:"postingDate,omitempty"`
}
