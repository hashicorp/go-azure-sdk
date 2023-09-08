package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelConfigurationOperationPredicate struct {
	DefaultDomainSuffix   *string
	Description           *string
	DisableUdpConnections *bool
	DisplayName           *string
	Id                    *string
	LastUpdateDateTime    *string
	ListenPort            *int64
	Network               *string
	ODataType             *string
}

func (p MicrosoftTunnelConfigurationOperationPredicate) Matches(input MicrosoftTunnelConfiguration) bool {

	if p.DefaultDomainSuffix != nil && (input.DefaultDomainSuffix == nil || *p.DefaultDomainSuffix != *input.DefaultDomainSuffix) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisableUdpConnections != nil && (input.DisableUdpConnections == nil || *p.DisableUdpConnections != *input.DisableUdpConnections) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastUpdateDateTime != nil && (input.LastUpdateDateTime == nil || *p.LastUpdateDateTime != *input.LastUpdateDateTime) {
		return false
	}

	if p.ListenPort != nil && (input.ListenPort == nil || *p.ListenPort != *input.ListenPort) {
		return false
	}

	if p.Network != nil && (input.Network == nil || *p.Network != *input.Network) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
