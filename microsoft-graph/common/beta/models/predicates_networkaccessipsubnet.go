package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIpSubnetOperationPredicate struct {
	ODataType *string
	Value     *string
}

func (p NetworkaccessIpSubnetOperationPredicate) Matches(input NetworkaccessIpSubnet) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
