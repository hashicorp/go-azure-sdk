package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureAdTokenAuthenticationOperationPredicate struct {
	ODataType  *string
	ResourceId *string
}

func (p AzureAdTokenAuthenticationOperationPredicate) Matches(input AzureAdTokenAuthentication) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	return true
}
