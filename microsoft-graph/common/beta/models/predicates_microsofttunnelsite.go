package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelSiteOperationPredicate struct {
	Description                     *string
	DisplayName                     *string
	Id                              *string
	InternalNetworkProbeUrl         *string
	ODataType                       *string
	PublicAddress                   *string
	UpgradeAutomatically            *bool
	UpgradeAvailable                *bool
	UpgradeWindowEndTime            *string
	UpgradeWindowStartTime          *string
	UpgradeWindowUtcOffsetInMinutes *int64
}

func (p MicrosoftTunnelSiteOperationPredicate) Matches(input MicrosoftTunnelSite) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InternalNetworkProbeUrl != nil && (input.InternalNetworkProbeUrl == nil || *p.InternalNetworkProbeUrl != *input.InternalNetworkProbeUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PublicAddress != nil && (input.PublicAddress == nil || *p.PublicAddress != *input.PublicAddress) {
		return false
	}

	if p.UpgradeAutomatically != nil && (input.UpgradeAutomatically == nil || *p.UpgradeAutomatically != *input.UpgradeAutomatically) {
		return false
	}

	if p.UpgradeAvailable != nil && (input.UpgradeAvailable == nil || *p.UpgradeAvailable != *input.UpgradeAvailable) {
		return false
	}

	if p.UpgradeWindowEndTime != nil && (input.UpgradeWindowEndTime == nil || *p.UpgradeWindowEndTime != *input.UpgradeWindowEndTime) {
		return false
	}

	if p.UpgradeWindowStartTime != nil && (input.UpgradeWindowStartTime == nil || *p.UpgradeWindowStartTime != *input.UpgradeWindowStartTime) {
		return false
	}

	if p.UpgradeWindowUtcOffsetInMinutes != nil && (input.UpgradeWindowUtcOffsetInMinutes == nil || *p.UpgradeWindowUtcOffsetInMinutes != *input.UpgradeWindowUtcOffsetInMinutes) {
		return false
	}

	return true
}
