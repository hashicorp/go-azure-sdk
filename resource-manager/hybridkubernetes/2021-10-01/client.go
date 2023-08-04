package v2021_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridkubernetes/2021-10-01/connectedclusters"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ConnectedClusters *connectedclusters.ConnectedClustersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	connectedClustersClient, err := connectedclusters.NewConnectedClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedClusters client: %+v", err)
	}
	configureFunc(connectedClustersClient.Client)

	return &Client{
		ConnectedClusters: connectedClustersClient,
	}, nil
}
