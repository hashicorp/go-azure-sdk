package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDnsCnameRecordOperationPredicate struct {
	CanonicalName    *string
	Id               *string
	IsOptional       *bool
	Label            *string
	ODataType        *string
	RecordType       *string
	SupportedService *string
	Ttl              *int64
}

func (p DomainDnsCnameRecordOperationPredicate) Matches(input DomainDnsCnameRecord) bool {

	if p.CanonicalName != nil && (input.CanonicalName == nil || *p.CanonicalName != *input.CanonicalName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsOptional != nil && (input.IsOptional == nil || *p.IsOptional != *input.IsOptional) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecordType != nil && (input.RecordType == nil || *p.RecordType != *input.RecordType) {
		return false
	}

	if p.SupportedService != nil && (input.SupportedService == nil || *p.SupportedService != *input.SupportedService) {
		return false
	}

	if p.Ttl != nil && (input.Ttl == nil || *p.Ttl != *input.Ttl) {
		return false
	}

	return true
}
