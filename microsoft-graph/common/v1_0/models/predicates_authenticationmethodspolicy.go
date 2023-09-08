package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicyOperationPredicate struct {
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	PolicyVersion        *string
	ReconfirmationInDays *int64
}

func (p AuthenticationMethodsPolicyOperationPredicate) Matches(input AuthenticationMethodsPolicy) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.PolicyVersion != nil && (input.PolicyVersion == nil || *p.PolicyVersion != *input.PolicyVersion) {
		return false
	}

	if p.ReconfirmationInDays != nil && (input.ReconfirmationInDays == nil || *p.ReconfirmationInDays != *input.ReconfirmationInDays) {
		return false
	}

	return true
}
