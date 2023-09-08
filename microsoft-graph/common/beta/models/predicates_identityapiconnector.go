package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityApiConnectorOperationPredicate struct {
	DisplayName *string
	Id          *string
	ODataType   *string
	TargetUrl   *string
}

func (p IdentityApiConnectorOperationPredicate) Matches(input IdentityApiConnector) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TargetUrl != nil && (input.TargetUrl == nil || *p.TargetUrl != *input.TargetUrl) {
		return false
	}

	return true
}
