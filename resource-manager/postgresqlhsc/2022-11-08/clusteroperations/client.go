package clusteroperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterOperationsClient struct {
	Client *resourcemanager.Client
}

func NewClusterOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ClusterOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "clusteroperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClusterOperationsClient: %+v", err)
	}

	return &ClusterOperationsClient{
		Client: client,
	}, nil
}
