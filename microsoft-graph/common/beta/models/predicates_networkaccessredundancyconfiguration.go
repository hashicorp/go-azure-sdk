package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRedundancyConfigurationOperationPredicate struct {
	ODataType          *string
	ZoneLocalIpAddress *string
}

func (p NetworkaccessRedundancyConfigurationOperationPredicate) Matches(input NetworkaccessRedundancyConfiguration) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ZoneLocalIpAddress != nil && (input.ZoneLocalIpAddress == nil || *p.ZoneLocalIpAddress != *input.ZoneLocalIpAddress) {
		return false
	}

	return true
}
