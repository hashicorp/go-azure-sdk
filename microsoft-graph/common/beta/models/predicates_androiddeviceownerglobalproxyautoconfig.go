package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGlobalProxyAutoConfigOperationPredicate struct {
	ODataType          *string
	ProxyAutoConfigURL *string
}

func (p AndroidDeviceOwnerGlobalProxyAutoConfigOperationPredicate) Matches(input AndroidDeviceOwnerGlobalProxyAutoConfig) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProxyAutoConfigURL != nil && (input.ProxyAutoConfigURL == nil || *p.ProxyAutoConfigURL != *input.ProxyAutoConfigURL) {
		return false
	}

	return true
}
