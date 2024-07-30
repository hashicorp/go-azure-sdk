package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConnectivityConfigurationLink struct {
	DisplayName         *string                                        `json:"displayName,omitempty"`
	Id                  *string                                        `json:"id,omitempty"`
	LocalConfigurations *[]NetworkaccessLocalConnectivityConfiguration `json:"localConfigurations,omitempty"`
	ODataType           *string                                        `json:"@odata.type,omitempty"`
	PeerConfiguration   *PeerConnectivityConfiguration                 `json:"peerConfiguration,omitempty"`
}
