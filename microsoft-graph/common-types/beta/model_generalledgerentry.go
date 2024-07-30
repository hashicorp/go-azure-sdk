package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeneralLedgerEntry struct {
	Account              *Account `json:"account,omitempty"`
	AccountId            *string  `json:"accountId,omitempty"`
	AccountNumber        *string  `json:"accountNumber,omitempty"`
	CreditAmount         *float64 `json:"creditAmount,omitempty"`
	DebitAmount          *float64 `json:"debitAmount,omitempty"`
	Description          *string  `json:"description,omitempty"`
	DocumentNumber       *string  `json:"documentNumber,omitempty"`
	DocumentType         *string  `json:"documentType,omitempty"`
	Id                   *string  `json:"id,omitempty"`
	LastModifiedDateTime *string  `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string  `json:"@odata.type,omitempty"`
	PostingDate          *string  `json:"postingDate,omitempty"`
}
