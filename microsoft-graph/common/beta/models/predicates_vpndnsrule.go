package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnDnsRuleOperationPredicate struct {
	AutoTrigger    *bool
	Name           *string
	ODataType      *string
	Persistent     *bool
	ProxyServerUri *string
}

func (p VpnDnsRuleOperationPredicate) Matches(input VpnDnsRule) bool {

	if p.AutoTrigger != nil && (input.AutoTrigger == nil || *p.AutoTrigger != *input.AutoTrigger) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Persistent != nil && (input.Persistent == nil || *p.Persistent != *input.Persistent) {
		return false
	}

	if p.ProxyServerUri != nil && (input.ProxyServerUri == nil || *p.ProxyServerUri != *input.ProxyServerUri) {
		return false
	}

	return true
}
