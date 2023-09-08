package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiApplicationOperationPredicate struct {
	AcceptMappedClaims          *bool
	ODataType                   *string
	RequestedAccessTokenVersion *int64
}

func (p ApiApplicationOperationPredicate) Matches(input ApiApplication) bool {

	if p.AcceptMappedClaims != nil && (input.AcceptMappedClaims == nil || *p.AcceptMappedClaims != *input.AcceptMappedClaims) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestedAccessTokenVersion != nil && (input.RequestedAccessTokenVersion == nil || *p.RequestedAccessTokenVersion != *input.RequestedAccessTokenVersion) {
		return false
	}

	return true
}
