package networking

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkingClient struct {
	Client *resourcemanager.Client
}

func NewNetworkingClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkingClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networking", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkingClient: %+v", err)
	}

	return &NetworkingClient{
		Client: client,
	}, nil
}
