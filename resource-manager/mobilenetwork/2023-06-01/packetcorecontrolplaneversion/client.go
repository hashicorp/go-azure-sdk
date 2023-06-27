package packetcorecontrolplaneversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPacketCoreControlPlaneVersionClientWithBaseURI(endpoint string) PacketCoreControlPlaneVersionClient {
	return PacketCoreControlPlaneVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
