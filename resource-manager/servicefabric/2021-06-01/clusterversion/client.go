package clusterversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterVersionClient struct {
	Client *resourcemanager.Client
}

func NewClusterVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterVersionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "clusterversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterVersionClient: %+v", err)
	}

	return &ClusterVersionClient{
		Client: client,
	}, nil
}
