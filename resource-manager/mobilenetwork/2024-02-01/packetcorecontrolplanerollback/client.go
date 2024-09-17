package packetcorecontrolplanerollback

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreControlPlaneRollbackClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreControlPlaneRollbackClientWithBaseURI(sdkApi sdkEnv.Api) (*PacketCoreControlPlaneRollbackClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "packetcorecontrolplanerollback", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreControlPlaneRollbackClient: %+v", err)
	}

	return &PacketCoreControlPlaneRollbackClient{
		Client: client,
	}, nil
}
