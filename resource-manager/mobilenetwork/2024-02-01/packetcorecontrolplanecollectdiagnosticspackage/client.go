package packetcorecontrolplanecollectdiagnosticspackage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneCollectDiagnosticsPackageClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlaneCollectDiagnosticsPackageClientWithBaseURI(sdkApi sdkEnv.Api) (*PacketCoreControlPlaneCollectDiagnosticsPackageClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "packetcorecontrolplanecollectdiagnosticspackage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlaneCollectDiagnosticsPackageClient: %+v", err)
	}

	return &PacketCoreControlPlaneCollectDiagnosticsPackageClient{
		Client: client,
	}, nil
}
