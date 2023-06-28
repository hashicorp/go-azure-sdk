package packetcaptures

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCapturesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPacketCapturesClientWithBaseURI(endpoint string) PacketCapturesClient {
	return PacketCapturesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
