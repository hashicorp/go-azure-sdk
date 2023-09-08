package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationOperationPredicate struct {
	ConnectAutomatically               *bool
	ConnectWhenNetworkNameIsHidden     *bool
	CreatedDateTime                    *string
	Description                        *string
	DisplayName                        *string
	Id                                 *string
	LastModifiedDateTime               *string
	NetworkName                        *string
	ODataType                          *string
	OuterIdentityPrivacyTemporaryValue *string
	PreSharedKey                       *string
	ProxyAutomaticConfigurationUrl     *string
	ProxyManualAddress                 *string
	ProxyManualPort                    *int64
	Ssid                               *string
	SupportsScopeTags                  *bool
	Version                            *int64
}

func (p MacOSEnterpriseWiFiConfigurationOperationPredicate) Matches(input MacOSEnterpriseWiFiConfiguration) bool {

	if p.ConnectAutomatically != nil && (input.ConnectAutomatically == nil || *p.ConnectAutomatically != *input.ConnectAutomatically) {
		return false
	}

	if p.ConnectWhenNetworkNameIsHidden != nil && (input.ConnectWhenNetworkNameIsHidden == nil || *p.ConnectWhenNetworkNameIsHidden != *input.ConnectWhenNetworkNameIsHidden) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.NetworkName != nil && (input.NetworkName == nil || *p.NetworkName != *input.NetworkName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OuterIdentityPrivacyTemporaryValue != nil && (input.OuterIdentityPrivacyTemporaryValue == nil || *p.OuterIdentityPrivacyTemporaryValue != *input.OuterIdentityPrivacyTemporaryValue) {
		return false
	}

	if p.PreSharedKey != nil && (input.PreSharedKey == nil || *p.PreSharedKey != *input.PreSharedKey) {
		return false
	}

	if p.ProxyAutomaticConfigurationUrl != nil && (input.ProxyAutomaticConfigurationUrl == nil || *p.ProxyAutomaticConfigurationUrl != *input.ProxyAutomaticConfigurationUrl) {
		return false
	}

	if p.ProxyManualAddress != nil && (input.ProxyManualAddress == nil || *p.ProxyManualAddress != *input.ProxyManualAddress) {
		return false
	}

	if p.ProxyManualPort != nil && (input.ProxyManualPort == nil || *p.ProxyManualPort != *input.ProxyManualPort) {
		return false
	}

	if p.Ssid != nil && (input.Ssid == nil || *p.Ssid != *input.Ssid) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
