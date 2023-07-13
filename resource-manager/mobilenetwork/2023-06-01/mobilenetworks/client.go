package mobilenetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileNetworksClient struct {
	Client *resourcemanager.Client
}

func NewMobileNetworksClientWithBaseURI(api environments.Api) (*MobileNetworksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "mobilenetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileNetworksClient: %+v", err)
	}

	return &MobileNetworksClient{
		Client: client,
	}, nil
}
