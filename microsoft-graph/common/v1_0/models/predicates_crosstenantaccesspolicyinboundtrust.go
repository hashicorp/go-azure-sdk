package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyInboundTrustOperationPredicate struct {
	IsCompliantDeviceAccepted           *bool
	IsHybridAzureADJoinedDeviceAccepted *bool
	IsMfaAccepted                       *bool
	ODataType                           *string
}

func (p CrossTenantAccessPolicyInboundTrustOperationPredicate) Matches(input CrossTenantAccessPolicyInboundTrust) bool {

	if p.IsCompliantDeviceAccepted != nil && (input.IsCompliantDeviceAccepted == nil || *p.IsCompliantDeviceAccepted != *input.IsCompliantDeviceAccepted) {
		return false
	}

	if p.IsHybridAzureADJoinedDeviceAccepted != nil && (input.IsHybridAzureADJoinedDeviceAccepted == nil || *p.IsHybridAzureADJoinedDeviceAccepted != *input.IsHybridAzureADJoinedDeviceAccepted) {
		return false
	}

	if p.IsMfaAccepted != nil && (input.IsMfaAccepted == nil || *p.IsMfaAccepted != *input.IsMfaAccepted) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
