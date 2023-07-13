package packetcorecontrolplaneversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneVersionClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlaneVersionClientWithBaseURI(api environments.Api) (*PacketCoreControlPlaneVersionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "packetcorecontrolplaneversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlaneVersionClient: %+v", err)
	}

	return &PacketCoreControlPlaneVersionClient{
		Client: client,
	}, nil
}
