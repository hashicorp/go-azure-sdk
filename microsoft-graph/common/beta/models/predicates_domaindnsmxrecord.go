package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainDnsMxRecordOperationPredicate struct {
	Id               *string
	IsOptional       *bool
	Label            *string
	MailExchange     *string
	ODataType        *string
	Preference       *int64
	RecordType       *string
	SupportedService *string
	Ttl              *int64
}

func (p DomainDnsMxRecordOperationPredicate) Matches(input DomainDnsMxRecord) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsOptional != nil && (input.IsOptional == nil || *p.IsOptional != *input.IsOptional) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.MailExchange != nil && (input.MailExchange == nil || *p.MailExchange != *input.MailExchange) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Preference != nil && (input.Preference == nil || *p.Preference != *input.Preference) {
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
