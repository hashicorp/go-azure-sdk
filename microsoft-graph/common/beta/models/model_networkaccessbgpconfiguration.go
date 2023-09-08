package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBgpConfiguration struct {
	Asn            *int64  `json:"asn,omitempty"`
	IpAddress      *string `json:"ipAddress,omitempty"`
	LocalIpAddress *string `json:"localIpAddress,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	PeerIpAddress  *string `json:"peerIpAddress,omitempty"`
}
