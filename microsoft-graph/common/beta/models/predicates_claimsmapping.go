package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsMappingOperationPredicate struct {
	DisplayName *string
	Email       *string
	GivenName   *string
	ODataType   *string
	Surname     *string
	UserId      *string
}

func (p ClaimsMappingOperationPredicate) Matches(input ClaimsMapping) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.GivenName != nil && (input.GivenName == nil || *p.GivenName != *input.GivenName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Surname != nil && (input.Surname == nil || *p.Surname != *input.Surname) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
