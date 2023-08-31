package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BgpServiceLoadBalancerConfiguration struct {
	BgpAdvertisements    *[]BgpAdvertisement           `json:"bgpAdvertisements,omitempty"`
	BgpPeers             *[]ServiceLoadBalancerBgpPeer `json:"bgpPeers,omitempty"`
	FabricPeeringEnabled *FabricPeeringEnabled         `json:"fabricPeeringEnabled,omitempty"`
	IPAddressPools       *[]IPAddressPool              `json:"ipAddressPools,omitempty"`
}
