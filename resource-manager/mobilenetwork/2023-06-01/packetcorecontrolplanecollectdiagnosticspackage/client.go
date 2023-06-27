package packetcorecontrolplanecollectdiagnosticspackage

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneCollectDiagnosticsPackageClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI(endpoint string) PacketCoreControlPlaneCollectDiagnosticsPackageClient {
	return PacketCoreControlPlaneCollectDiagnosticsPackageClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
