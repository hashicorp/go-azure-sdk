package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContractOperationPredicate struct {
	ContractType      *string
	CustomerId        *string
	DefaultDomainName *string
	DeletedDateTime   *string
	DisplayName       *string
	Id                *string
	ODataType         *string
}

func (p ContractOperationPredicate) Matches(input Contract) bool {

	if p.ContractType != nil && (input.ContractType == nil || *p.ContractType != *input.ContractType) {
		return false
	}

	if p.CustomerId != nil && (input.CustomerId == nil || *p.CustomerId != *input.CustomerId) {
		return false
	}

	if p.DefaultDomainName != nil && (input.DefaultDomainName == nil || *p.DefaultDomainName != *input.DefaultDomainName) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
