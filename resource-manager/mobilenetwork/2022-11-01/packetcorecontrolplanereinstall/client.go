package packetcorecontrolplanereinstall

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneReinstallClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlaneReinstallClientWithBaseURI(api environments.Api) (*PacketCoreControlPlaneReinstallClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "packetcorecontrolplanereinstall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlaneReinstallClient: %+v", err)
	}

	return &PacketCoreControlPlaneReinstallClient{
		Client: client,
	}, nil
}
