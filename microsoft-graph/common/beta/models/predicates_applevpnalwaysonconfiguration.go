package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationOperationPredicate struct {
	AllowAllCaptiveNetworkPlugins *bool
	AllowCaptiveWebSheet          *bool
	NatKeepAliveIntervalInSeconds *int64
	NatKeepAliveOffloadEnable     *bool
	ODataType                     *string
	UserToggleEnabled             *bool
}

func (p AppleVpnAlwaysOnConfigurationOperationPredicate) Matches(input AppleVpnAlwaysOnConfiguration) bool {

	if p.AllowAllCaptiveNetworkPlugins != nil && (input.AllowAllCaptiveNetworkPlugins == nil || *p.AllowAllCaptiveNetworkPlugins != *input.AllowAllCaptiveNetworkPlugins) {
		return false
	}

	if p.AllowCaptiveWebSheet != nil && (input.AllowCaptiveWebSheet == nil || *p.AllowCaptiveWebSheet != *input.AllowCaptiveWebSheet) {
		return false
	}

	if p.NatKeepAliveIntervalInSeconds != nil && (input.NatKeepAliveIntervalInSeconds == nil || *p.NatKeepAliveIntervalInSeconds != *input.NatKeepAliveIntervalInSeconds) {
		return false
	}

	if p.NatKeepAliveOffloadEnable != nil && (input.NatKeepAliveOffloadEnable == nil || *p.NatKeepAliveOffloadEnable != *input.NatKeepAliveOffloadEnable) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserToggleEnabled != nil && (input.UserToggleEnabled == nil || *p.UserToggleEnabled != *input.UserToggleEnabled) {
		return false
	}

	return true
}
