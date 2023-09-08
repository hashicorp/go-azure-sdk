package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfaceOperationPredicate struct {
	Description      *string
	IpV4Address      *string
	IpV6Address      *string
	LocalIpV6Address *string
	MacAddress       *string
	ODataType        *string
}

func (p NetworkInterfaceOperationPredicate) Matches(input NetworkInterface) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.IpV4Address != nil && (input.IpV4Address == nil || *p.IpV4Address != *input.IpV4Address) {
		return false
	}

	if p.IpV6Address != nil && (input.IpV6Address == nil || *p.IpV6Address != *input.IpV6Address) {
		return false
	}

	if p.LocalIpV6Address != nil && (input.LocalIpV6Address == nil || *p.LocalIpV6Address != *input.LocalIpV6Address) {
		return false
	}

	if p.MacAddress != nil && (input.MacAddress == nil || *p.MacAddress != *input.MacAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
