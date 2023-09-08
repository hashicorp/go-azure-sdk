package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GeneralLedgerEntryOperationPredicate struct {
	AccountId            *string
	AccountNumber        *string
	CreditAmount         *float64
	DebitAmount          *float64
	Description          *string
	DocumentNumber       *string
	DocumentType         *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	PostingDate          *string
}

func (p GeneralLedgerEntryOperationPredicate) Matches(input GeneralLedgerEntry) bool {

	if p.AccountId != nil && (input.AccountId == nil || *p.AccountId != *input.AccountId) {
		return false
	}

	if p.AccountNumber != nil && (input.AccountNumber == nil || *p.AccountNumber != *input.AccountNumber) {
		return false
	}

	if p.CreditAmount != nil && (input.CreditAmount == nil || *p.CreditAmount != *input.CreditAmount) {
		return false
	}

	if p.DebitAmount != nil && (input.DebitAmount == nil || *p.DebitAmount != *input.DebitAmount) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DocumentNumber != nil && (input.DocumentNumber == nil || *p.DocumentNumber != *input.DocumentNumber) {
		return false
	}

	if p.DocumentType != nil && (input.DocumentType == nil || *p.DocumentType != *input.DocumentType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PostingDate != nil && (input.PostingDate == nil || *p.PostingDate != *input.PostingDate) {
		return false
	}

	return true
}
