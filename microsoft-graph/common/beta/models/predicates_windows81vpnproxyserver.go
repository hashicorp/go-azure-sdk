package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81VpnProxyServerOperationPredicate struct {
	Address                          *string
	AutomaticConfigurationScriptUrl  *string
	AutomaticallyDetectProxySettings *bool
	BypassProxyServerForLocalAddress *bool
	ODataType                        *string
	Port                             *int64
}

func (p Windows81VpnProxyServerOperationPredicate) Matches(input Windows81VpnProxyServer) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.AutomaticConfigurationScriptUrl != nil && (input.AutomaticConfigurationScriptUrl == nil || *p.AutomaticConfigurationScriptUrl != *input.AutomaticConfigurationScriptUrl) {
		return false
	}

	if p.AutomaticallyDetectProxySettings != nil && (input.AutomaticallyDetectProxySettings == nil || *p.AutomaticallyDetectProxySettings != *input.AutomaticallyDetectProxySettings) {
		return false
	}

	if p.BypassProxyServerForLocalAddress != nil && (input.BypassProxyServerForLocalAddress == nil || *p.BypassProxyServerForLocalAddress != *input.BypassProxyServerForLocalAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Port != nil && (input.Port == nil || *p.Port != *input.Port) {
		return false
	}

	return true
}
