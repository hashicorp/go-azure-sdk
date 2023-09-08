package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OAuthConsentAppDetailOperationPredicate struct {
	DisplayLogo *string
	DisplayName *string
	ODataType   *string
}

func (p OAuthConsentAppDetailOperationPredicate) Matches(input OAuthConsentAppDetail) bool {

	if p.DisplayLogo != nil && (input.DisplayLogo == nil || *p.DisplayLogo != *input.DisplayLogo) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
