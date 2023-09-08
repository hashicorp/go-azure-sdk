package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainIdentitySourceOperationPredicate struct {
	DisplayName *string
	DomainName  *string
	ODataType   *string
}

func (p DomainIdentitySourceOperationPredicate) Matches(input DomainIdentitySource) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DomainName != nil && (input.DomainName == nil || *p.DomainName != *input.DomainName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
