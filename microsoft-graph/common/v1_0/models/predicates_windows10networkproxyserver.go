package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10NetworkProxyServerOperationPredicate struct {
	Address              *string
	ODataType            *string
	UseForLocalAddresses *bool
}

func (p Windows10NetworkProxyServerOperationPredicate) Matches(input Windows10NetworkProxyServer) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UseForLocalAddresses != nil && (input.UseForLocalAddresses == nil || *p.UseForLocalAddresses != *input.UseForLocalAddresses) {
		return false
	}

	return true
}
