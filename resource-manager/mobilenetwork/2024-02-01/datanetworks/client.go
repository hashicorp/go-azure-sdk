package datanetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataNetworksClient struct {
	Client *resourcemanager.Client
}

func NewDataNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*DataNetworksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "datanetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataNetworksClient: %+v", err)
	}

	return &DataNetworksClient{
		Client: client,
	}, nil
}
