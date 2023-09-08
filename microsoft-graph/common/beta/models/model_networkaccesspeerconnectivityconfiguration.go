package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPeerConnectivityConfiguration struct {
	Asn        *int64  `json:"asn,omitempty"`
	BgpAddress *string `json:"bgpAddress,omitempty"`
	Endpoint   *string `json:"endpoint,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
}
