package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostSecurityStateOperationPredicate struct {
	Fqdn                      *string
	IsAzureAdJoined           *bool
	IsAzureAdRegistered       *bool
	IsHybridAzureDomainJoined *bool
	NetBiosName               *string
	ODataType                 *string
	Os                        *string
	PrivateIpAddress          *string
	PublicIpAddress           *string
	RiskScore                 *string
}

func (p HostSecurityStateOperationPredicate) Matches(input HostSecurityState) bool {

	if p.Fqdn != nil && (input.Fqdn == nil || *p.Fqdn != *input.Fqdn) {
		return false
	}

	if p.IsAzureAdJoined != nil && (input.IsAzureAdJoined == nil || *p.IsAzureAdJoined != *input.IsAzureAdJoined) {
		return false
	}

	if p.IsAzureAdRegistered != nil && (input.IsAzureAdRegistered == nil || *p.IsAzureAdRegistered != *input.IsAzureAdRegistered) {
		return false
	}

	if p.IsHybridAzureDomainJoined != nil && (input.IsHybridAzureDomainJoined == nil || *p.IsHybridAzureDomainJoined != *input.IsHybridAzureDomainJoined) {
		return false
	}

	if p.NetBiosName != nil && (input.NetBiosName == nil || *p.NetBiosName != *input.NetBiosName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Os != nil && (input.Os == nil || *p.Os != *input.Os) {
		return false
	}

	if p.PrivateIpAddress != nil && (input.PrivateIpAddress == nil || *p.PrivateIpAddress != *input.PrivateIpAddress) {
		return false
	}

	if p.PublicIpAddress != nil && (input.PublicIpAddress == nil || *p.PublicIpAddress != *input.PublicIpAddress) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	return true
}
