package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUserOperationPredicate struct {
	AccountDomain     *string
	AccountName       *string
	FirstSeenDateTime *string
	LastSeenDateTime  *string
	LogonId           *string
	ODataType         *string
}

func (p LogonUserOperationPredicate) Matches(input LogonUser) bool {

	if p.AccountDomain != nil && (input.AccountDomain == nil || *p.AccountDomain != *input.AccountDomain) {
		return false
	}

	if p.AccountName != nil && (input.AccountName == nil || *p.AccountName != *input.AccountName) {
		return false
	}

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
		return false
	}

	if p.LogonId != nil && (input.LogonId == nil || *p.LogonId != *input.LogonId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
