package provisionednetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedNetworksClient struct {
	Client *resourcemanager.Client
}

func NewProvisionedNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*ProvisionedNetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "provisionednetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProvisionedNetworksClient: %+v", err)
	}

	return &ProvisionedNetworksClient{
		Client: client,
	}, nil
}
