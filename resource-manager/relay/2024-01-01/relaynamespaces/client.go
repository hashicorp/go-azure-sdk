package relaynamespaces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelayNamespacesClient struct {
	Client *resourcemanager.Client
}

func NewRelayNamespacesClientWithBaseURI(sdkApi sdkEnv.Api) (*RelayNamespacesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "relaynamespaces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RelayNamespacesClient: %+v", err)
	}

	return &RelayNamespacesClient{
		Client: client,
	}, nil
}
