package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedClaimConfigurationOperationPredicate struct {
	NameIdPolicyFormat *string
	ODataType          *string
}

func (p SupportedClaimConfigurationOperationPredicate) Matches(input SupportedClaimConfiguration) bool {

	if p.NameIdPolicyFormat != nil && (input.NameIdPolicyFormat == nil || *p.NameIdPolicyFormat != *input.NameIdPolicyFormat) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
