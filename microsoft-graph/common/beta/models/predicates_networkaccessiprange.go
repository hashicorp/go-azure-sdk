package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIpRangeOperationPredicate struct {
	BeginAddress *string
	EndAddress   *string
	ODataType    *string
}

func (p NetworkaccessIpRangeOperationPredicate) Matches(input NetworkaccessIpRange) bool {

	if p.BeginAddress != nil && (input.BeginAddress == nil || *p.BeginAddress != *input.BeginAddress) {
		return false
	}

	if p.EndAddress != nil && (input.EndAddress == nil || *p.EndAddress != *input.EndAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
