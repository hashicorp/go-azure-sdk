package hypervcluster

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVClusterClient struct {
	Client *resourcemanager.Client
}

func NewHyperVClusterClientWithBaseURI(sdkApi sdkEnv.Api) (*HyperVClusterClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hypervcluster", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HyperVClusterClient: %+v", err)
	}

	return &HyperVClusterClient{
		Client: client,
	}, nil
}
