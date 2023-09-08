package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationOperationPredicate struct {
	ODataType    *string
	PreSharedKey *string
}

func (p NetworkaccessTunnelConfigurationOperationPredicate) Matches(input NetworkaccessTunnelConfiguration) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PreSharedKey != nil && (input.PreSharedKey == nil || *p.PreSharedKey != *input.PreSharedKey) {
		return false
	}

	return true
}
