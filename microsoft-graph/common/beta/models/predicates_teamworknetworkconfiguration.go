package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkNetworkConfigurationOperationPredicate struct {
	DefaultGateway  *string
	DomainName      *string
	HostName        *string
	IpAddress       *string
	IsDhcpEnabled   *bool
	IsPCPortEnabled *bool
	ODataType       *string
	PrimaryDns      *string
	SecondaryDns    *string
	SubnetMask      *string
}

func (p TeamworkNetworkConfigurationOperationPredicate) Matches(input TeamworkNetworkConfiguration) bool {

	if p.DefaultGateway != nil && (input.DefaultGateway == nil || *p.DefaultGateway != *input.DefaultGateway) {
		return false
	}

	if p.DomainName != nil && (input.DomainName == nil || *p.DomainName != *input.DomainName) {
		return false
	}

	if p.HostName != nil && (input.HostName == nil || *p.HostName != *input.HostName) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.IsDhcpEnabled != nil && (input.IsDhcpEnabled == nil || *p.IsDhcpEnabled != *input.IsDhcpEnabled) {
		return false
	}

	if p.IsPCPortEnabled != nil && (input.IsPCPortEnabled == nil || *p.IsPCPortEnabled != *input.IsPCPortEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrimaryDns != nil && (input.PrimaryDns == nil || *p.PrimaryDns != *input.PrimaryDns) {
		return false
	}

	if p.SecondaryDns != nil && (input.SecondaryDns == nil || *p.SecondaryDns != *input.SecondaryDns) {
		return false
	}

	if p.SubnetMask != nil && (input.SubnetMask == nil || *p.SubnetMask != *input.SubnetMask) {
		return false
	}

	return true
}
