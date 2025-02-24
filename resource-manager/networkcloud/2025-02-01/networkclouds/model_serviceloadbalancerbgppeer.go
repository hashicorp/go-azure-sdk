package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceLoadBalancerBgpPeer struct {
	BfdEnabled    *BfdEnabled  `json:"bfdEnabled,omitempty"`
	BgpMultiHop   *BgpMultiHop `json:"bgpMultiHop,omitempty"`
	HoldTime      *string      `json:"holdTime,omitempty"`
	KeepAliveTime *string      `json:"keepAliveTime,omitempty"`
	MyAsn         *int64       `json:"myAsn,omitempty"`
	Name          string       `json:"name"`
	Password      *string      `json:"password,omitempty"`
	PeerAddress   string       `json:"peerAddress"`
	PeerAsn       int64        `json:"peerAsn"`
	PeerPort      *int64       `json:"peerPort,omitempty"`
}
