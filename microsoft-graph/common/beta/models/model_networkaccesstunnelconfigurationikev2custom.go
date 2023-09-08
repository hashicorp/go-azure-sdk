package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTunnelConfigurationIKEv2Custom struct {
	DhGroup           *NetworkaccessTunnelConfigurationIKEv2CustomDhGroup         `json:"dhGroup,omitempty"`
	IkeEncryption     *NetworkaccessTunnelConfigurationIKEv2CustomIkeEncryption   `json:"ikeEncryption,omitempty"`
	IkeIntegrity      *NetworkaccessTunnelConfigurationIKEv2CustomIkeIntegrity    `json:"ikeIntegrity,omitempty"`
	IpSecEncryption   *NetworkaccessTunnelConfigurationIKEv2CustomIpSecEncryption `json:"ipSecEncryption,omitempty"`
	IpSecIntegrity    *NetworkaccessTunnelConfigurationIKEv2CustomIpSecIntegrity  `json:"ipSecIntegrity,omitempty"`
	ODataType         *string                                                     `json:"@odata.type,omitempty"`
	PfsGroup          *NetworkaccessTunnelConfigurationIKEv2CustomPfsGroup        `json:"pfsGroup,omitempty"`
	PreSharedKey      *string                                                     `json:"preSharedKey,omitempty"`
	SaLifeTimeSeconds *int64                                                      `json:"saLifeTimeSeconds,omitempty"`
}
