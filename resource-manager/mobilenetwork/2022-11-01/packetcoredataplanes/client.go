package packetcoredataplanes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreDataPlanesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPacketCoreDataPlanesClientWithBaseURI(endpoint string) PacketCoreDataPlanesClient {
	return PacketCoreDataPlanesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
