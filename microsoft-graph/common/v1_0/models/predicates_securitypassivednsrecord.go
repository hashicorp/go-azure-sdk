package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPassiveDnsRecordOperationPredicate struct {
	CollectedDateTime *string
	FirstSeenDateTime *string
	Id                *string
	LastSeenDateTime  *string
	ODataType         *string
	RecordType        *string
}

func (p SecurityPassiveDnsRecordOperationPredicate) Matches(input SecurityPassiveDnsRecord) bool {

	if p.CollectedDateTime != nil && (input.CollectedDateTime == nil || *p.CollectedDateTime != *input.CollectedDateTime) {
		return false
	}

	if p.FirstSeenDateTime != nil && (input.FirstSeenDateTime == nil || *p.FirstSeenDateTime != *input.FirstSeenDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSeenDateTime != nil && (input.LastSeenDateTime == nil || *p.LastSeenDateTime != *input.LastSeenDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecordType != nil && (input.RecordType == nil || *p.RecordType != *input.RecordType) {
		return false
	}

	return true
}
