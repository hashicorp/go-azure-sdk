package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBgpConfigurationOperationPredicate struct {
	Asn            *int64
	IpAddress      *string
	LocalIpAddress *string
	ODataType      *string
	PeerIpAddress  *string
}

func (p NetworkaccessBgpConfigurationOperationPredicate) Matches(input NetworkaccessBgpConfiguration) bool {

	if p.Asn != nil && (input.Asn == nil || *p.Asn != *input.Asn) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.LocalIpAddress != nil && (input.LocalIpAddress == nil || *p.LocalIpAddress != *input.LocalIpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PeerIpAddress != nil && (input.PeerIpAddress == nil || *p.PeerIpAddress != *input.PeerIpAddress) {
		return false
	}

	return true
}
