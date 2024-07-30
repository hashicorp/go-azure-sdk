package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyInboundTrust struct {
	IsCompliantDeviceAccepted           *bool   `json:"isCompliantDeviceAccepted,omitempty"`
	IsHybridAzureADJoinedDeviceAccepted *bool   `json:"isHybridAzureADJoinedDeviceAccepted,omitempty"`
	IsMfaAccepted                       *bool   `json:"isMfaAccepted,omitempty"`
	ODataType                           *string `json:"@odata.type,omitempty"`
}
