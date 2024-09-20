package packetcoredataplanes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreDataPlanesClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreDataPlanesClientWithBaseURI(sdkApi sdkEnv.Api) (*PacketCoreDataPlanesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "packetcoredataplanes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreDataPlanesClient: %+v", err)
	}

	return &PacketCoreDataPlanesClient{
		Client: client,
	}, nil
}
